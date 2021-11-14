package mq

import (
	"fmt"
	"liveearth/infrastructure/config"
	"time"

	"github.com/beanstalkd/go-beanstalk"

	"github.com/sereiner/library/concurrent/cmap"
)

type Mq interface {
	// topic 队列名称
	// body 消息体
	// pri 优先级
	// delay 延迟
	// ttr 超时
	Send(topic string, body []byte, pri uint32, delay, ttr time.Duration) (id uint64, err error)
	Conn() *beanstalk.Conn
	// Delete 删除任务
	// id   job id
	Delete(id uint64) error
}

type IComponentMQ interface {
	GetRegularMQ(names ...string) (d Mq)
	GetMQ(names ...string) (d Mq, err error)
	GetMQBy(name string) (c Mq, err error)
	SaveMQObject(name string, f func(conf config.MqConfig) (Mq, error)) (bool, Mq, error)
	Close() error
}

type StandardMQ struct {
	name  string
	conn  *beanstalk.Conn
	tMap  map[string]*beanstalk.Tube
	mqMap cmap.ConcurrentMap
}

func NewStandardMQ(name ...string) IComponentMQ {
	if len(name) > 0 {
		return &StandardMQ{
			name:  name[0],
			tMap:  make(map[string]*beanstalk.Tube),
			mqMap: cmap.New(2),
		}
	}
	return &StandardMQ{
		name:  "default",
		tMap:  make(map[string]*beanstalk.Tube),
		mqMap: cmap.New(2),
	}
}

func (s *StandardMQ) Send(topic string, body []byte, pri uint32, delay, ttr time.Duration) (id uint64, err error) {

	v, ok := s.tMap[topic]
	if !ok {
		v = beanstalk.NewTube(s.conn, topic)
		s.tMap[topic] = v
	}

	return v.Put(body, pri, delay, ttr)
}

func (s *StandardMQ) Delete(id uint64) error {
	return s.conn.Delete(id)
}

func (s *StandardMQ) Conn() *beanstalk.Conn {
	return s.conn
}

func (s *StandardMQ) GetRegularMQ(names ...string) (d Mq) {
	d, err := s.GetMQ(names...)
	if err != nil {
		panic(err)
	}

	return d
}

func (s *StandardMQ) GetMQ(names ...string) (d Mq, err error) {
	name := s.name
	if len(names) > 0 {
		name = names[0]
	}
	return s.GetMQBy(name)
}

func (s *StandardMQ) GetMQBy(name string) (c Mq, err error) {

	_, c, err = s.SaveMQObject(name, func(conf config.MqConfig) (Mq, error) {

		if s.conn == nil {
			conn, err := beanstalk.Dial("tcp", conf.Address)
			if err != nil {
				return nil, err
			}
			s.conn = conn
		}

		return s, nil
	})

	return c, err
}

func (s *StandardMQ) SaveMQObject(name string, f func(conf config.MqConfig) (Mq, error)) (bool, Mq, error) {

	key := fmt.Sprintf("%s/%s", "mq", name)

	ok, ch, err := s.mqMap.SetIfAbsentCb(key, func(input ...interface{}) (c interface{}, err error) {

		mqConf, ok := config.C.MQ[name]
		if !ok {
			panic(fmt.Sprintf("数据库配置不存在 name:%s", name))
		}

		return f(mqConf)
	})
	if err != nil {
		err = fmt.Errorf("创建mq失败 err:%v", err)
		return ok, nil, err
	}

	return ok, ch.(Mq), err
}

func (s *StandardMQ) Close() error {
	s.mqMap.RemoveIterCb(func(k string, v interface{}) bool {

		return true
	})
	return nil
}
