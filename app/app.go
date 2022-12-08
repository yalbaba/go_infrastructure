package app

import (
	"flag"
	"fmt"
	"github.com/yalbaba/go_infrastructure/servers/ws"

	"github.com/yalbaba/go_infrastructure/component/registry"

	"github.com/yalbaba/go_infrastructure/consts"
	"github.com/yalbaba/go_infrastructure/pkg/logfilter"
	"github.com/yalbaba/go_infrastructure/servers/cron"
	"github.com/yalbaba/go_infrastructure/servers/http"
	"github.com/yalbaba/go_infrastructure/servers/mqc"
	"github.com/yalbaba/go_infrastructure/servers/nsq_consume"

	logger "github.com/sereiner/library/log"

	"github.com/mikegleasonjr/workers"

	"sync"

	"github.com/yalbaba/go_infrastructure/servers"

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

	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/config"
	_ "github.com/yalbaba/go_infrastructure/servers/cron"
	_ "github.com/yalbaba/go_infrastructure/servers/http"
	_ "github.com/yalbaba/go_infrastructure/servers/mqc"
	_ "github.com/yalbaba/go_infrastructure/servers/nsq_consume"
	_ "github.com/yalbaba/go_infrastructure/servers/rpc"
	_ "github.com/yalbaba/go_infrastructure/servers/ws"

	"github.com/yalbaba/go_infrastructure/pkg/iris"

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
	RegisterWs(path string, handler ws.Handler)
	GetContainer() component.Container

	Run() (string, error)
	Close()
	logger.ILogger
}

type GApp struct {
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

func NewGApp(opts ...Option) IApp {

	initConfig()

	server := &GApp{
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

	flag.StringVar(&configPath, "f", "", "server config path")
	flag.Parse()

	if configPath == "" {
		// log.Warn("未指定配置文件路径! 将使用 ./configs/config_dev.toml 配置文件加载程序")

		configPath = envs.GetString("CACHE_CONFIG", "./server_configs/config_dev.toml")
	}

	if _, err := toml.DecodeFile(configPath, &config.C); err != nil {
		log.Panic(err)
	}

}

func (s *GApp) GetContainer() component.Container {
	return s.component
}

func (s *GApp) RegisterAPIRouter(f func(component.Container, iris.Party)) {

	if s.servers[consts.HttpServer] == nil {
		return
	}

	s.servers[consts.HttpServer].RegisterService(http.RegisterAPIFunc(f))

}

func (s *GApp) RegisterMqcWorker(topic string, handler workers.HandlerFunc) {

	if s.servers[consts.MqcServer] == nil {
		return
	}

	s.servers[consts.MqcServer].RegisterService(mqc.MqcHandlers{topic: handler})
}

func (s *GApp) RegisterCronJob(name string, cronStr string, disable bool, handler cron.Handler) {
	if s.servers[consts.CronServer] == nil {
		return
	}

	s.servers[consts.CronServer].RegisterService(cron.NewCronTask(&cron.Task{
		Name:    name,
		Cron:    cronStr,
		Disable: disable,
		Handler: handler}, s.Logger))
}

func (s *GApp) RegisterRpcService(sc ...interface{}) {

	if s.servers[consts.RpcServer] == nil {
		return
	}

	s.servers[consts.RpcServer].RegisterService(sc...)
}

func (s *GApp) RegisterNsqHandler(topic, channel string, handler nsq_consume.RegistryNsqConsumerHandlerFunc, opts ...nsq_consume.ConsumerOption) {
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

func (s *GApp) RegisterMidJob(f func(component.Container)) {
	s.midJobs = append(s.midJobs, f)
}

func (s *GApp) start() {

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

func (s *GApp) RegisterWs(path string, handler ws.Handler) {
	if handler == nil || len(path) == 0 {
		return
	}

	s.servers[consts.WsServer].RegisterService(map[string]ws.Handler{
		path: handler,
	})
}

func (s *GApp) freeMemory() {
	for {
		select {
		case <-s.closeChan:
			return
		case <-time.After(time.Second * 120):
			debug.FreeOSMemory()
		}
	}
}

func (s *GApp) Close() {
	s.done = true
	close(s.closeChan)
	s.interrupt <- syscall.SIGTERM
	for _, v := range s.servers {
		_ = v.Close()
	}
	s.component.Close()

}

func (s *GApp) Run() (string, error) {

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
