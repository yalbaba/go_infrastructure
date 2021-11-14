package websocket

import "liveearth/infrastructure/servers/websocket/conn"

type HandlerFunc func(data conn.Event, conn conn.IConn) interface{}

type Engine struct {
	addr string
	*RouterGroup
	groups []*RouterGroup
}

func NewEngine() *Engine {
	e := &Engine{
		groups: make([]*RouterGroup, 0),
	}
	e.RouterGroup = &RouterGroup{engine: e} //继承RouterGroup的方法
	return e
}

type IGroup interface {
	Group(p string) IGroup
	RegisterEvent(event string, handler HandlerFunc)
	RegisterController(c ConnController)
}

type RouterGroup struct {
	path       string
	engine     *Engine
	router     *Router
	controller ConnController
}

func (rg *RouterGroup) Group(p string) IGroup {
	engine := rg.engine
	newGroup := &RouterGroup{
		engine: engine,
		path:   rg.path + p,
		router: newRouter(),
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (rg *RouterGroup) RegisterEvent(event string, handler HandlerFunc) {
	rg.router.handlers[event] = handler
}

func (rg *RouterGroup) RegisterController(c ConnController) {
	rg.controller = c
}

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) addEvent(event string, handler HandlerFunc) {
	r.handlers[event] = handler
}
