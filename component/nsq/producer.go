/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/3/1
   Description :
-------------------------------------------------
*/

package nsq

import (
	"fmt"
	"time"

	nnsq "github.com/nsqio/go-nsq"

	"liveearth/infrastructure/component/nsq/conn"
	"liveearth/infrastructure/config"
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
)

type IComponentNsq interface {
	GetNsq(names ...string) *nnsq.Producer
	// 关闭
	Close()
}

type instance struct {
	producer *nnsq.Producer
}

func (i *instance) Close() {
	i.producer.Stop()
}

type NsqProducer struct {
	conn *conn.Conn
}

func NewNsqProducer() IComponentNsq {
	n := &NsqProducer{
		conn: conn.NewConn(),
	}
	return n
}

func (r *NsqProducer) GetNsq(name ...string) *nnsq.Producer {
	return r.conn.GetInstance(r.makeClient, name...).(*instance).producer
}

func (r *NsqProducer) makeClient(name string) (conn.IInstance, error) {
	conf, ok := config.C.Nsq[name]
	if !ok {
		return nil, fmt.Errorf("组件配置<nsq.%s>不存在", name)
	}

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
	if conf.Address == "" {
		return nil, fmt.Errorf("nsq的address为空")
	}

	nsqConf := nnsq.NewConfig()
	nsqConf.AuthSecret = conf.AuthSecret
	nsqConf.HeartbeatInterval = time.Duration(conf.HeartbeatInterval) * time.Millisecond
	nsqConf.ReadTimeout = time.Duration(conf.ReadTimeout) * time.Millisecond
	nsqConf.WriteTimeout = time.Duration(conf.WriteTimeout) * time.Millisecond
	nsqConf.DialTimeout = time.Duration(conf.DialTimeout) * time.Millisecond

	producer, err := nnsq.NewProducer(conf.Address, nsqConf)
	return &instance{producer}, err
}

func (r *NsqProducer) Close() {
	r.conn.CloseAll()
}
