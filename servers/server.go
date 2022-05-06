package servers

import (
	"fmt"
	"go_infrastructure/component"
	"go_infrastructure/consts"
)

type IServer interface {
	RegisterService(sc ...interface{})
	Start() error
	Close() error
	GetServerType() consts.ServerType
}

type IServerResolver interface {
	Resolve(container component.Container) IServer
}

type IServerResolverHandler func(container component.Container) IServer

var resolvers = make(map[consts.ServerType]IServerResolver)

func Register(serverType consts.ServerType, resolver IServerResolver) {

	if _, ok := resolvers[serverType]; ok {
		panic("server: Register called twice for identifier: " + serverType.String())
	}
	resolvers[serverType] = resolver
}

func NewServer(serverType consts.ServerType, c component.Container) IServer {

	if resolver, ok := resolvers[serverType]; ok {
		return resolver.Resolve(c)
	}
	panic(fmt.Errorf("server: unknown identifier name %q (forgotten import?)", serverType.String()))
}
