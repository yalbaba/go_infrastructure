package beanstalk

import (
	"sync"
	"time"

	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/config"
	"github.com/yalbaba/go_infrastructure/consts"
	"github.com/yalbaba/go_infrastructure/servers"

	"github.com/mikegleasonjr/workers"
)

type MqcHandlers map[string]workers.HandlerFunc

type BeanstalkServer struct {
	workMux *workers.WorkMux
	c       component.Container
	lock    sync.Mutex
}

func (m *BeanstalkServer) GetServerType() consts.ServerType {
	return consts.BeanstalkServer
}

func NewBeanstalkServer(c component.Container) *BeanstalkServer {
	return &BeanstalkServer{
		workMux: workers.NewWorkMux(),
		c:       c,
	}
}

func (m *BeanstalkServer) RegisterService(sc ...interface{}) {

	jm, ok := sc[0].(MqcHandlers)
	if !ok {
		panic("mqc 注册函数错误")
	}

	for k, v := range jm {
		m.lock.Lock()
		m.workMux.Handle(k, v)
		m.lock.Unlock()
	}
}

func (m *BeanstalkServer) Start() error {

	m.c.Debug("开始启动 MQC 服务器...")
	errChan2 := make(chan error, 1)
	go func(errChan2 chan error) {
		for _, topic := range m.workMux.Tubes() {
			m.c.Debug("开始监听 " + topic)
		}
		if err := workers.ConnectAndWork("tcp", config.C.MQ["default"].Address, m.workMux); err != nil {
			errChan2 <- err
		}
	}(errChan2)

	select {
	case <-time.After(time.Millisecond * 500):
	case err := <-errChan2:
		m.c.Error(err.Error())
	}

	m.c.Debug("MQC 服务器启动成功")
	return nil
}

func (m *BeanstalkServer) Close() error {
	return nil
}

type BeanstalkServerAdapter struct {
}

func (h *BeanstalkServerAdapter) Resolve(c component.Container) servers.IServer {
	return NewBeanstalkServer(c)
}

func init() {
	servers.Register(consts.BeanstalkServer, &BeanstalkServerAdapter{})
}
