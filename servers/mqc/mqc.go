package mqc

import (
	"go_infrastructure/component"
	"go_infrastructure/config"
	"go_infrastructure/consts"
	"go_infrastructure/servers"
	"sync"
	"time"

	"github.com/mikegleasonjr/workers"
)

type MqcHandlers map[string]workers.HandlerFunc

type MqcServer struct {
	workMux *workers.WorkMux
	c       component.Container
	lock    sync.Mutex
}

func (m *MqcServer) GetServerType() consts.ServerType {
	return consts.MqcServer
}

func NewMqcServer(c component.Container) *MqcServer {
	return &MqcServer{
		workMux: workers.NewWorkMux(),
		c:       c,
	}
}

func (m *MqcServer) RegisterService(sc ...interface{}) {

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

func (m *MqcServer) Start() error {

	m.c.Debug("开始启动 MQC 服务器...")
	errChan2 := make(chan error, 1)
	go func(errChan2 chan error) {
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

func (m *MqcServer) Close() error {
	return nil
}

type mqcServerAdapter struct {
}

func (h *mqcServerAdapter) Resolve(c component.Container) servers.IServer {
	return NewMqcServer(c)
}

func init() {
	servers.Register(consts.MqcServer, &mqcServerAdapter{})
}
