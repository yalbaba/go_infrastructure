/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/1/25
   Description :
-------------------------------------------------
*/

package nsq_consume

import (
	"fmt"
	"strings"
	"time"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"

	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/config"
	"github.com/yalbaba/go_infrastructure/utils"
)

type Context struct {
	*nsq.Message
	Topic               string
	Channel             string
	disableAutoRequeued bool // 关闭自动重排
}

// 关闭自动重排
func (ctx *Context) DisableAutoRequeued() {
	ctx.disableAutoRequeued = true
}

type RegistryNsqConsumerHandlerFunc = func(ctx *Context) error

type ConsumerConfig struct {
	Topic   string
	Channel string
	Handler RegistryNsqConsumerHandlerFunc
	Opts    []ConsumerOption
	*config.NsqConsumeConfig
}

type consumerCli struct {
	c        component.Container
	conf     *ConsumerConfig
	consumer *nsq.Consumer
	*consumerOptions
}

func newConsumer(c component.Container, conf *ConsumerConfig) *consumerCli {
	consumer := &consumerCli{
		c:               c,
		conf:            conf,
		consumerOptions: newConsumerOptions(),
	}

	for _, o := range conf.Opts {
		o(consumer.consumerOptions)
	}

	if consumer.ConsumeAttempts == 0 {
		consumer.ConsumeAttempts = conf.ConsumeAttempts
	}
	return consumer
}

func (c *consumerCli) Start() error {
	if c.Disable {
		return nil
	}

	// 构建配置
	nsqConf := nsq.NewConfig()
	nsqConf.AuthSecret = c.conf.AuthSecret
	nsqConf.HeartbeatInterval = time.Duration(c.conf.HeartbeatInterval) * time.Millisecond
	nsqConf.ReadTimeout = time.Duration(c.conf.ReadTimeout) * time.Millisecond
	nsqConf.WriteTimeout = time.Duration(c.conf.WriteTimeout) * time.Millisecond
	nsqConf.DialTimeout = time.Duration(c.conf.DialTimeout) * time.Millisecond
	nsqConf.DefaultRequeueDelay = time.Duration(c.conf.RequeueDelay) * time.Millisecond
	nsqConf.MaxRequeueDelay = time.Duration(c.conf.MaxRequeueDelay) * time.Millisecond
	nsqConf.MaxInFlight = c.conf.MaxInFlight
	nsqConf.MaxAttempts = 0

	// 创建消费者
	consumer, err := nsq.NewConsumer(c.conf.Topic, c.conf.Channel, nsqConf)
	if err != nil {
		return fmt.Errorf("创建nsq消费者失败, topic:%s, channel:%s, err:%s", c.conf.Topic, c.conf.Channel, err)
	}
	c.consumer = consumer

	// 添加消费handler
	threadCount := c.ThreadCount
	if threadCount == 0 { // 如果为0使用默认配置
		threadCount = c.conf.ThreadCount
	}
	c.consumer.AddConcurrentHandlers(c, threadCount)

	// 连接
	if c.conf.NsqLookupdAddress != "" {
		addresses := strings.Split(c.conf.NsqLookupdAddress, ",")
		return c.consumer.ConnectToNSQLookupds(addresses)
	}
	addresses := strings.Split(c.conf.Address, ",")
	return c.consumer.ConnectToNSQDs(addresses)
}

func (c *consumerCli) Close() error {
	if c.Disable {
		return nil
	}

	c.consumer.Stop()
	<-c.consumer.StopChan
	return nil
}

func (c *consumerCli) HandleMessage(message *nsq.Message) error {
	ctx := &Context{
		Message: message,
		Topic:   c.conf.Topic,
		Channel: c.conf.Channel,
	}

	c.c.Debug("nsqConsumer.receive")
	err := utils.Recover.WrapCall(func() error {
		return c.conf.Handler(ctx)
	})

	if err == nil {
		c.c.Debug("nsqConsumer.success")
		return nil
	}

	// 如果关闭了自动重排
	if ctx.disableAutoRequeued {
		c.c.Error("nsqConsumer.error! and requeued is closed", zap.Error(err))
		return nil
	}

	// 检查自动重排次数
	if ctx.Attempts >= c.ConsumeAttempts {
		c.c.Error("nsqConsumer.error! reach the maximum automatic Requeue Attempts", zap.Error(err))
		return nil
	}

	c.c.Error("nsqConsumer.error!", zap.Error(err))
	return err
}
