package http

import (
	"context"
	"liveearth/infrastructure/component"
	"liveearth/infrastructure/config"
	"liveearth/infrastructure/consts"
	"liveearth/infrastructure/pkg/iris"
	"liveearth/infrastructure/servers"
	"liveearth/infrastructure/servers/http/middleware"
	"liveearth/infrastructure/servers/http/middleware/cors"
	"time"

	inet "github.com/sereiner/library/net"
)

type RegisterAPIFunc func(component.Container, iris.Party)

type HttpServer struct {
	*iris.Application
	c component.Container
}

func (h *HttpServer) GetServerType() consts.ServerType {
	return consts.HttpServer
}

func NewHttpServer(c component.Container) *HttpServer {

	irisAPI := iris.New()
	irisAPI.Logger().SetLevel("disable")
	irisAPI.Use(
		middleware.CheckAuth(),
		middleware.LoggerMiddleware(),
		cors.AllowAll(),
		middleware.Recover(c),
	)
	irisAPI.AllowMethods(iris.MethodOptions)

	return &HttpServer{Application: irisAPI, c: c}
}

func (h *HttpServer) RegisterService(sc ...interface{}) {

	if len(sc) != 1 {
		h.c.Error("http 服务注册函数错误")
		return
	}

	f, ok := sc[0].(RegisterAPIFunc)
	if !ok {
		h.c.Error("http 服务注册函数类型错误")
		return
	}

	f(h.c, h.Party("/"))

}

func (h *HttpServer) Start() error {

	h.c.Debug("开始启动 API 服务器...")
	errChan2 := make(chan error, 1)
	go func(errChan2 chan error) {
		var addr string
		if config.C.Debug {
			addr = config.C.API.Addr
		} else {
			addr = inet.GetLocalIPAddress() + config.C.API.Addr
		}

		if err := h.Run(iris.Addr(addr),
			iris.WithoutBodyConsumptionOnUnmarshal,
			iris.WithoutPathCorrection,
			iris.WithOptimizations,
			iris.WithRemoteAddrHeader("X-Real-IP"),
		); err != nil && err != iris.ErrServerClosed {
			errChan2 <- err
		}
	}(errChan2)

	select {
	case <-time.After(time.Millisecond * 500):
	case err := <-errChan2:
		h.c.Error(err.Error())
	}

	h.c.Debugf("API 服务器启动成功 addr->[ %s ]", config.C.API.Addr)
	return nil
}

func (h *HttpServer) Close() error {
	return h.Shutdown(context.Background())
}

type httpServerAdapter struct {
}

func (h *httpServerAdapter) Resolve(c component.Container) servers.IServer {
	return NewHttpServer(c)
}

func init() {
	servers.Register(consts.HttpServer, &httpServerAdapter{})
}
