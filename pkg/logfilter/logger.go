package logfilter

import (
	"fmt"
	"strings"
	"sync"
	"time"

	logger "github.com/sereiner/library/log"
)

type loggerSetting struct {
	Level    string `json:"level" valid:"in(Off|Debug|Info|Warn|Error|Fatal|All),required"`
	Service  string `json:"service" valid:"required"`
	Interval string `json:"interval" valid:"required"`
}

type RPCLogger struct {
	platName    string
	systemName  string
	serverTypes []string
	clusterName string
	logger      *logger.Logger
	writer      *rpcWriter
	appenders   []*RPCAppender
	service     string
	appender    *logger.Appender
	//currentConf *conf.JSONConf
	closeChan chan struct{}
	once      sync.Once
	lock      sync.RWMutex
}

//NewRPCLogger 创建RPC日志程序
func NewRPCLogger(log *logger.Logger, platName string, systemName string, clusterName string, serverTypes []string) (r *RPCLogger, err error) {
	r = &RPCLogger{
		platName:    platName,
		systemName:  systemName,
		clusterName: clusterName,
		serverTypes: serverTypes,
		closeChan:   make(chan struct{}),
		logger:      log,
		appenders:   make([]*RPCAppender, 0, 2),
		appender:    &logger.Appender{Type: "rpc", Level: "Info", Interval: "@every 1m"},
	}

	go r.loopWatch()
	return r, r.changed()
}

func (r *RPCLogger) loopWatch() {
	tkr := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-r.closeChan:
			return
		case <-tkr.C:
			tkr.Stop()
		}
	}
}

//MakeAppender 构建Appender
func (r *RPCLogger) MakeAppender(l *logger.Appender, event *logger.LogEvent) (logger.IAppender, error) {

	r.lock.RLock()
	defer r.lock.RUnlock()

	iAppender, err := NewRPCAppender(r.writer, r.appender)
	if err != nil {
		return nil, err
	}

	r.appenders = append(r.appenders, iAppender)

	return iAppender, nil
}

//GetType 日志类型
func (r *RPCLogger) GetType() string {
	return "rpc"
}

//MakeUniq 获取日志标识
func (r *RPCLogger) MakeUniq(l *logger.Appender, event *logger.LogEvent) string {
	return "rpc"
}

func (r *RPCLogger) changed() error {

	var setting = loggerSetting{
		Level:    "Error",
		Service:  "",
		Interval: "1s",
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	if r.service != setting.Service {
		_, domain, server, err := resolvePath(setting.Service, "", "")
		if err != nil || domain == "" || server == "" {
			return fmt.Errorf("%s不合法 %v", setting.Service, err)
		}

		r.service = setting.Service
	}

	writer := newRPCWriter(setting.Service, r.platName, r.systemName, r.clusterName, r.serverTypes)
	r.writer = writer

	r.appender.Type = "rpc"
	r.appender.Level = setting.Level
	r.appender.Layout = `{"time":"%datetime","content":"%content","level":"%l","session":"%session"}`
	r.appender.Interval = setting.Interval

	for _, app := range r.appenders {
		app.Reset(setting.Interval, writer)
	}

	r.once.Do(func() {
		logger.RegistryFactory(r, r.appender)
	})

	return nil
}

//Close 关闭RPC日志
func (r *RPCLogger) Close() error {
	close(r.closeChan)
	return nil
}

//解析日志
func resolvePath(address string, d string, s string) (service string, domain string, server string, err error) {
	raddress := strings.TrimRight(address, "@")
	addrs := strings.SplitN(raddress, "@", 2)
	if len(addrs) == 1 {
		if addrs[0] == "" {
			return "", "", "", fmt.Errorf("服务地址%s不能为空", address)
		}
		service = "/" + strings.Trim(strings.Replace(raddress, ".", "/", -1), "/")
		domain = d
		server = s
		return
	}
	if addrs[0] == "" {
		return "", "", "", fmt.Errorf("%s错误，服务名不能为空", address)
	}
	if addrs[1] == "" {
		return "", "", "", fmt.Errorf("%s错误，服务名，域不能为空", address)
	}
	service = "/" + strings.Trim(strings.Replace(addrs[0], ".", "/", -1), "/")
	raddr := strings.Split(strings.TrimRight(addrs[1], "."), ".")
	if len(raddr) >= 2 && raddr[0] != "" && raddr[1] != "" {
		domain = raddr[len(raddr)-1]
		server = strings.Join(raddr[0:len(raddr)-1], ".")
		return
	}
	if len(raddr) == 1 {
		if raddr[0] == "" {
			return "", "", "", fmt.Errorf("%s错误，服务器名称不能为空", address)
		}
		domain = d
		server = raddr[0]
		return
	}
	if raddr[0] == "" && raddr[1] == "" {
		return "", "", "", fmt.Errorf(`%s错误,未指定服务器名称和域名称`, addrs[1])
	}
	domain = raddr[1]
	server = s
	return
}
