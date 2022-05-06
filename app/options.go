package app

import "liveearth/infrastructure/consts"

type option struct {
	PlatName    string
	AppName     string
	ServerTypes map[consts.ServerType]bool
	IsDebug     bool
	Trace       string
	HasRegistry bool
}

type Option func(*option)

func WithPlatName(name string) Option {
	return func(o *option) {
		o.PlatName = name
	}
}
func WithAppName(name string) Option {
	return func(o *option) {
		o.AppName = name
	}
}

func WithAPI() Option {
	return func(o *option) {
		o.ServerTypes[consts.HttpServer] = true
	}
}

func WithGRPC() Option {
	return func(o *option) {
		o.ServerTypes[consts.RpcServer] = true
	}
}

func WithMQC() Option {
	return func(o *option) {
		o.ServerTypes[consts.MqcServer] = true
	}
}

func WithCron() Option {
	return func(o *option) {
		o.ServerTypes[consts.CronServer] = true
	}
}

func WithNsqConsume() Option {
	return func(o *option) {
		o.ServerTypes[consts.NsqConsumeServer] = true
	}
}

func WithRegistry() Option {
	return func(o *option) {
		o.HasRegistry = true
	}
}

func WithWs() Option {
	return func(o *option) {
		o.ServerTypes[consts.WsServer] = true
	}
}
