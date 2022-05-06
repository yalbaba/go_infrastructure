package app

import (
	"flag"
	"fmt"
	"liveearth/infrastructure/component/registry"

	"liveearth/infrastructure/consts"
	"liveearth/infrastructure/pkg/logfilter"
	"liveearth/infrastructure/servers/cron"
	"liveearth/infrastructure/servers/http"
	"liveearth/infrastructure/servers/mqc"
	"liveearth/infrastructure/servers/nsq_consume"

	logger "github.com/sereiner/library/log"

	"github.com/mikegleasonjr/workers"

	"sync"

	"liveearth/infrastructure/servers"

	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/sereiner/library/envs"
	"github.com/takama/daemon"
	"github.com/wule61/log"
	"go.uber.org/zap"

	"time"

	"liveearth/infrastructure/component"
	"liveearth/infrastructure/config"
	_ "liveearth/infrastructure/servers/cron"
	_ "liveearth/infrastructure/servers/http"
	_ "liveearth/infrastructure/servers/mqc"
	_ "liveearth/infrastructure/servers/nsq_consume"
	_ "liveearth/infrastructure/servers/rpc"
	_ "liveearth/infrastructure/servers/ws"

	"liveearth/infrastructure/pkg/iris"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var configPath string

type IApp interface {
	RegisterRpcService(...interface{})
	RegisterAPIRouter(func(component.Container, iris.Party))
	RegisterMqcWorker(topic string, handler workers.HandlerFunc)
	RegisterCronJob(name string, cron string, disable bool, handler cron.Handler)
	RegisterNsqHandler(topic, channel string, handler nsq_consume.RegistryNsqConsumerHandlerFunc, opts ...nsq_consume.ConsumerOption)
	RegisterMidJob(f func(component.Container))
	RegisterWs(path string, handler http.Handler)
	GetContainer() component.Container

	Run() (string, error)
	Close()
	logger.ILogger
}

type UranusApp struct {
	component component.IComponent
	servers   map[consts.ServerType]servers.IServer
	daemon.Daemon
	*option
	*logger.Logger
	closeChan chan struct{}
	interrupt chan os.Signal
	done      bool
	mux       sync.Mutex
	midJobs   []func(c component.Container)
}

func NewUranusApp(opts ...Option) IApp {

	initConfig()

	server := &UranusApp{
		closeChan: make(chan struct{}),
		interrupt: make(chan os.Signal, 1),
		option:    &option{ServerTypes: make(map[consts.ServerType]bool)},
		servers:   make(map[consts.ServerType]servers.IServer),
	}

	for _, opt := range opts {
		opt(server.option)
	}

	dm, err := daemon.New(server.AppName, server.AppName)
	if err != nil {
		panic(err)
	}
	server.Daemon = dm

	server.Logger = logger.GetSession("system", logger.CreateSession())

	c := component.NewComponent(server.Logger)
	if server.HasRegistry {
		c.IRegistry = registry.NewRegistry(server.Logger)
	}
	server.component = c

	_, err = logfilter.NewRPCLogger(server.Logger, "live", server.AppName, "c", []string{})
	if err != nil {
		panic(err)
	}

	for k, v := range server.ServerTypes {
		server.mux.Lock()
		if v {
			server.servers[k] = servers.NewServer(k, c)
		}
		server.mux.Unlock()
	}

	return server

}

func initConfig() {

	flag.StringVar(&configPath, "c", "", "grpc_server config path")
	flag.Parse()

	if configPath == "" {
		// log.Warn("未指定配置文件路径! 将使用 ./configs/config_dev.toml 配置文件加载程序")

		configPath = envs.GetString("CACHE_CONFIG", "./configs/config_dev.toml")
	}

	if _, err := toml.DecodeFile(configPath, &config.C); err != nil {
		log.Panic(err)
	}

}

func (s *UranusApp) GetContainer() component.Container {
	return s.component
}

func (s *UranusApp) RegisterAPIRouter(f func(component.Container, iris.Party)) {

	if s.servers[consts.HttpServer] == nil {
		return
	}

	s.servers[consts.HttpServer].RegisterService(http.RegisterAPIFunc(f))

}

func (s *UranusApp) RegisterMqcWorker(topic string, handler workers.HandlerFunc) {

	if s.servers[consts.MqcServer] == nil {
		return
	}

	s.servers[consts.MqcServer].RegisterService(mqc.MqcHandlers{topic: handler})
}

func (s *UranusApp) RegisterCronJob(name string, cronStr string, disable bool, handler cron.Handler) {
	if s.servers[consts.CronServer] == nil {
		return
	}

	s.servers[consts.CronServer].RegisterService(cron.NewCronTask(&cron.Task{
		Name:    name,
		Cron:    cronStr,
		Disable: disable,
		Handler: handler}, s.Logger))
}

func (s *UranusApp) RegisterRpcService(sc ...interface{}) {

	if s.servers[consts.RpcServer] == nil {
		return
	}

	s.servers[consts.RpcServer].RegisterService(sc...)
}

func (s *UranusApp) RegisterNsqHandler(topic, channel string, handler nsq_consume.RegistryNsqConsumerHandlerFunc, opts ...nsq_consume.ConsumerOption) {
	if s.servers[consts.NsqConsumeServer] == nil {
		return
	}

	s.servers[consts.NsqConsumeServer].RegisterService(&nsq_consume.ConsumerConfig{
		Topic:   topic,
		Channel: channel,
		Handler: handler,
		Opts:    opts,
	})
}

func (s *UranusApp) RegisterMidJob(f func(component.Container)) {
	s.midJobs = append(s.midJobs, f)
}

func (s *UranusApp) start() {

	for _, v := range s.servers {
		if err := v.Start(); err != nil {
			panic(fmt.Errorf("启动服务器失败, 服务器类型:%s err:%v", v.GetServerType().String(), zap.Error(err)))
		}
	}

	if s.HasRegistry {
		s.component.Debug("正在注册服务...")
		if err := s.component.GetRegistry().Register(); err != nil {
			panic(err)
		}
		s.component.Debugf("成功注册所有服务 services->  ")
		for _, v := range config.C.RegisterServerList {
			s.component.Debug("service_name: ", v.ServiceName, " | service_info: ", v.ServerInfo, " | ttl: ", v.TTl)
		}
	}

	for i := 0; i < len(s.midJobs); i++ {
		go s.midJobs[i](s.GetContainer())
	}

	go s.freeMemory()

	signal.Notify(s.interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
LOOP:
	for {
		select {
		case <-s.interrupt:
			s.done = true
			break LOOP
		}
	}
	s.Debug(fmt.Sprintf("%s 服务器正在退出...", s.AppName))
	s.Close()
	s.Debug(fmt.Sprintf("%s 服务器已安全退出...", s.AppName))
	time.Sleep(time.Second)
}

func (s *UranusApp) RegisterWs(path string, handler http.Handler) {
	if handler == nil || len(path) == 0 {
		return
	}

	s.servers[consts.WsServer].RegisterService(map[string]http.Handler{
		path: handler,
	})
}

func (s *UranusApp) freeMemory() {
	for {
		select {
		case <-s.closeChan:
			return
		case <-time.After(time.Second * 120):
			debug.FreeOSMemory()
		}
	}
}

func (s *UranusApp) Close() {
	s.done = true
	close(s.closeChan)
	s.interrupt <- syscall.SIGTERM
	for _, v := range s.servers {
		_ = v.Close()
	}
	s.component.Close()

}

func (s *UranusApp) Run() (string, error) {

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return s.Install(os.Args[2:]...)
		case "remove":
			return s.Remove()
		case "start":
			return s.Start()
		case "stop":
			return s.Stop()
		case "status":
			return s.Status()
		}
	}

	s.start()
	return "", nil
}
