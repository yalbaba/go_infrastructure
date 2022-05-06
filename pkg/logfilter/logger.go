package logfilter

import (
	"fmt"
	"sync"
	"time"

	logger "github.com/sereiner/library/log"
	"github.com/sereiner/parrot/conf"
	"github.com/sereiner/parrot/rpc"
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
	currentConf *conf.JSONConf
	closeChan   chan struct{}
	once        sync.Once
	lock        sync.RWMutex
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
		_, domain, server, err := rpc.ResolvePath(setting.Service, "", "")
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
