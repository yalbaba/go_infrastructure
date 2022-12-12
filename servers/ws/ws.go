package ws

import (
	sctx "context"
	inet "github.com/sereiner/library/net"
	http2 "net/http"
	"time"

	"github.com/yalbaba/go_infrastructure/component"
	"github.com/yalbaba/go_infrastructure/config"
	"github.com/yalbaba/go_infrastructure/consts"
	"github.com/yalbaba/go_infrastructure/pkg/exchange"
	"github.com/yalbaba/go_infrastructure/pkg/iris"
	"github.com/yalbaba/go_infrastructure/servers"
	"github.com/yalbaba/go_infrastructure/servers/http/middleware"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	c          component.Container
	dispatcher *iris.Application
	closeChan  chan struct{}
	send       chan []byte
}

func NewWsServer(c component.Container) *WsServer {

	w := &WsServer{
		c:         c,
		closeChan: make(chan struct{}),
		send:      make(chan []byte, 256),
	}

	w.dispatcher = iris.New()
	w.dispatcher.Use(
		middleware.Recover(c),
		middleware.CheckAuth(),
		middleware.LoggerMiddleware("ws"),
	)

	return w
}

func (w *WsServer) RegisterService(sc ...interface{}) {
	if len(sc) != 1 {
		w.c.Error("ws 服务注册函数错误")
		return
	}

	m, ok := sc[0].(map[string]Handler)
	if !ok {
		w.c.Error("ws 服务注册函数类型错误")
		return
	}

	for k, v := range m {
		w.dispatcher.Any(k, w.handler(v))
	}

}

var (
	upgrader = websocket.Upgrader{}
)

func (w *WsServer) handler(hf Handler) iris.Handler {

	return func(ctx iris.Context) {

		upgrader.CheckOrigin = func(r *http2.Request) bool {
			return true
		}

		ws, err := upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
		if err != nil {
			w.c.Error(err)
			return
		}

		h := newWSHandler(ws)
		exchange.WSExchange.Subscribe(getUUID(ctx), h.recvNotify(ctx))
		defer exchange.WSExchange.Unsubscribe(getUUID(ctx))

		go h.readPump(hf, ctx)
		h.writePump()
	}
}

func (w *WsServer) Start() error {

	w.c.Debug("开始启动 WS 服务器")
	errChan := make(chan error, 1)
	go func(ch chan error) {
		var addr string
		if config.C.Debug {
			addr = config.C.API.Addr
		} else {
			addr = inet.GetLocalIPAddress() + config.C.WS.Addr
		}

		if err := w.dispatcher.Run(iris.Addr(addr),
			iris.WithoutBodyConsumptionOnUnmarshal,
			iris.WithoutPathCorrection,
			iris.WithOptimizations,
			iris.WithRemoteAddrHeader("X-Real-IP"),
		); err != nil && err != iris.ErrServerClosed {
			errChan <- err
		}
	}(errChan)
	select {
	case <-time.After(time.Millisecond * 500):

	case err := <-errChan:
		w.c.Error(err)
	}

	w.c.Debugf("WS 服务器启动成功 addr->[ %s ]", config.C.WS.Addr)
	return nil

}

func (w *WsServer) Close() error {
	return w.dispatcher.Shutdown(sctx.Background())
}

func (w *WsServer) GetServerType() consts.ServerType {
	return consts.WsServer
}

type wsServerAdapter struct {
}

func (h *wsServerAdapter) Resolve(c component.Container) servers.IServer {
	return NewWsServer(c)
}

func init() {
	servers.Register(consts.WsServer, &wsServerAdapter{})
}
