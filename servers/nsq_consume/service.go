/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/3/1
   Description :
-------------------------------------------------
*/

package nsq_consume

import (
	"runtime"
	"sync"

	"go_infrastructure/component"
	"go_infrastructure/config"
	"go_infrastructure/consts"
	"go_infrastructure/servers"
)

const (
	// 默认心跳间隔
	defaultHeartbeatInterval = 30000
	// 默认读取超时
	defaultReadTimeout = 30000
	// 默认写入超时
	defaultWriteTimeout = 5000
	// 默认连接超时
	defaultDialTimeout = 2000
	// MaxInFlight
	defaultMaxInFlight = 1024
	// 默认延时时间
	defaultRequeueDelay = 60000
	// 默认最大延时时间
	defaultMaxRequeueDelay = 600000
	// 默认消费尝试次数
	defaultConsumeAttempts = 3
)

type nsqConsumerServerAdapter struct{}

func (h *nsqConsumerServerAdapter) Resolve(c component.Container) servers.IServer {
	return NewNsqConsumerServer(c)
}

func init() {
	servers.Register(consts.NsqConsumeServer, &nsqConsumerServerAdapter{})
}

type NsqConsumeServer struct {
	c         component.Container
	conf      *config.NsqConsumeConfig
	consumers []*consumerCli
}

func (n *NsqConsumeServer) RegisterService(sc ...interface{}) {
	for _, v := range sc {
		conf, ok := v.(*ConsumerConfig)
		if !ok {
			n.c.Fatal("nsq消费服务注入类型错误, 它必须能转为 *nsqc.ConsumerConfig")
		}
		conf.NsqConsumeConfig = n.conf

		consumer := newConsumer(n.c, conf)
		n.consumers = append(n.consumers, consumer)
	}
}

func (n *NsqConsumeServer) Start() error {
	// 开始消费
	for _, consumer := range n.consumers {
		if err := consumer.Start(); err != nil {
			return err
		}
	}

	return nil
}

func (n *NsqConsumeServer) Close() error {
	var wg sync.WaitGroup
	wg.Add(len(n.consumers))
	for _, consumer := range n.consumers {
		go func(consumer *consumerCli) {
			defer wg.Done()
			_ = consumer.Close()
		}(consumer)
	}
	wg.Wait()
	return nil
}

func (n *NsqConsumeServer) GetServerType() consts.ServerType {
	return consts.NsqConsumeServer
}

func NewNsqConsumerServer(c component.Container) *NsqConsumeServer {
	conf := config.C.NsqConsume

	if conf.ReadTimeout <= 0 {
		conf.ReadTimeout = defaultReadTimeout
	}
	if conf.WriteTimeout <= 0 {
		conf.WriteTimeout = defaultWriteTimeout
	}
	if conf.DialTimeout <= 0 {
		conf.DialTimeout = defaultDialTimeout
	}
	if conf.HeartbeatInterval <= 0 {
		conf.HeartbeatInterval = defaultHeartbeatInterval
	}
	if conf.HeartbeatInterval > conf.ReadTimeout {
		conf.HeartbeatInterval = conf.ReadTimeout
	}
	if conf.MaxInFlight <= 0 {
		conf.MaxInFlight = defaultMaxInFlight
	}
	if conf.ThreadCount <= 0 {
		conf.ThreadCount = runtime.NumCPU()
	}
	if conf.RequeueDelay <= 0 {
		conf.RequeueDelay = defaultRequeueDelay
	}
	if conf.MaxRequeueDelay <= 0 {
		conf.MaxRequeueDelay = defaultMaxRequeueDelay
	}
	if conf.ConsumeAttempts == 0 {
		conf.ConsumeAttempts = defaultConsumeAttempts
	}

	if conf.Address == "" && conf.NsqLookupdAddress == "" {
		c.Fatal("nsq_consume的address为空")
	}

	return &NsqConsumeServer{
		c:    c,
		conf: &conf,
	}
}
