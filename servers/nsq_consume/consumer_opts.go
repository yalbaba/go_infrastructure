/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/3/1
   Description :
-------------------------------------------------
*/

package nsq_consume

type consumerOptions struct {
	Disable         bool // 禁用
	ThreadCount     int
	ConsumeAttempts uint16
}

type ConsumerOption func(opts *consumerOptions)

func newConsumerOptions() *consumerOptions {
	return &consumerOptions{
		Disable:         false,
		ThreadCount:     0,
		ConsumeAttempts: 0,
	}
}

// 禁用
func WithConsumerDisable(disable ...bool) ConsumerOption {
	return func(opts *consumerOptions) {
		opts.Disable = len(disable) == 0 || disable[0]
	}
}

// 线程数, 默认为0表示使用配置的默认线程数
//
// 同时处理信息的goroutine数
func WithConsumerThreadCount(threadCount int) ConsumerOption {
	return func(opts *consumerOptions) {
		if threadCount < 0 {
			threadCount = 0
		}
		opts.ThreadCount = threadCount
	}
}

// 消费尝试次数, 默认为0表示使用全局配置的次数
func WithConsumerAttempts(attempts uint16) ConsumerOption {
	return func(opts *consumerOptions) {
		opts.ConsumeAttempts = attempts
	}
}
