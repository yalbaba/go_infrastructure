package websocket

import (
	"github.com/fasthttp-contrib/websocket"
	gw "github.com/gorilla/websocket"
	inet "github.com/sereiner/library/net"
	"github.com/sereiner/library/types"
	"liveearth/infrastructure/component"
	"liveearth/infrastructure/config"
	"liveearth/infrastructure/consts"
	"liveearth/infrastructure/servers"
	"liveearth/infrastructure/servers/websocket/conn"
	"liveearth/infrastructure/utils"
	"net"
	"net/http"
	"time"
)

type RegisterWsFunc func(component.Container, IGroup)

type WebSocketServer struct {
	engine *Engine
	c      component.Container
}

func NewWebSocketServer(c component.Container) *WebSocketServer {
	return &WebSocketServer{
		engine: NewEngine(),
		c:      c,
	}
}

func (s *WebSocketServer) RegisterService(sc ...interface{}) {

	if len(sc) != 1 {
		s.c.Error("ws 服务注册函数错误")
		return
	}

	f, ok := sc[0].(RegisterWsFunc)
	if !ok {
		s.c.Error("ws 服务注册函数类型错误")
		return
	}

	f(s.c, s.engine.Group(""))
}

func (s *WebSocketServer) Start() error {

	var addr string
	if config.C.Debug {
		addr = config.C.WsApi.Addr
	} else {
		addr = inet.GetLocalIPAddress() + config.C.WsApi.Addr
	}

	var upgrader = gw.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				return false
			}
			return true
		},
	}

	for i, group := range s.engine.groups {
		if i == 0 || group.controller == nil {
			continue
		}
		f := s.build(upgrader, group)
		http.HandleFunc(group.path, f)
	}

	errChan := make(chan error, 1)
	s.c.Debug("开始启动 Websocket 服务器...")
	go func(ec chan error) {
		if err := http.ListenAndServe(addr, nil); err != nil {
			ec <- err
		}
	}(errChan)

	s.c.Debugf("Websocket 服务器启动成功 addr->[ %s ]", addr)
	select {
	case <-time.After(time.Millisecond * 500):
	case err := <-errChan:
		s.c.Error(err.Error())
	}

	return nil
}

func (s *WebSocketServer) Close() error {
	return nil
}

func (s *WebSocketServer) GetServerType() consts.ServerType {
	return consts.WebSocketServer
}

func (s *WebSocketServer) build(upgrader gw.Upgrader, group *RouterGroup) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			s.c.Errorf("连接失败,err:%v", err)
			return
		}

		connObj := conn.NewConn(utils.ProduceUserIdByUuid(),
			types.GetInt64(r.URL.Query()["group_id"][0]),
			types.GetString(r.URL.Query()["device_id"][0]),
			types.GetInt(r.URL.Query()["platform"][0]),
			wsConn)

		s.c.SaveConn(connObj)

		res, err := group.controller.OnConnect(connObj)
		if err != nil {
			s.c.Errorf("OnConnect err:%v", err)
			return
		}

		err = connObj.Send("Connect", res)
		if err != nil {
			s.c.Errorf("Send Response err:%v", err)
			return
		}

		s.c.Infof("Session: [%s], Event: Connect, Response: %v", connObj.GetDeviceId(), string(conn.WriteResponse("Connect", res).Data.([]byte)))

		defer func() {
			err := group.controller.OnClose(connObj)
			if err != nil {
				s.c.Errorf("OnClose err:%v", err)
			}

			err = connObj.Close()
			if err != nil {
				s.c.Errorf("Close err:%v", err)
			}

			s.c.DeleteConn(connObj.GetDeviceId())

		}()

		for {
			defer func() {
				if err := recover(); err != nil {
					s.c.Errorf("panic err: %v", err)
				}
			}()

			_, message, err := connObj.GetConn().ReadMessage()
			if err != nil {
				if len(message) == 0 {
					continue
				}
				if netErr, ok := err.(net.Error); ok {
					if netErr.Timeout() {
						break
					}
				}
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
					s.c.Errorf("ReadMessage other remote:%v error: %v", connObj.GetConn().RemoteAddr(), err)
					break
				}
			}

			request, err := conn.WriteRequest(message)
			if err != nil {
				s.c.Errorf("WriteRequest 转换参数失败,err: %v, source: %s", err, string(message))
				continue
			}

			s.c.Infof("Session: [%s], Event: %s, Request: %v", connObj.GetDeviceId(), request.EventName, request)

			res := group.router.handlers[request.EventName](request, connObj)

			err = connObj.Send(request.EventName, res)
			if err != nil {
				s.c.Errorf("响应执行结果失败,err: %v", err)
				continue
			}
			s.c.Infof("Session: [%s], Event: %s, Response: %v", connObj.GetDeviceId(), request.EventName, string(conn.WriteResponse(request.EventName, res).Data.([]byte)))
		}
	}

}

/*
适配器
*/
type websocketServerAdapter struct {
}

func (w *websocketServerAdapter) Resolve(container component.Container) servers.IServer {
	return NewWebSocketServer(container)
}

func init() {
	servers.Register(consts.WebSocketServer, &websocketServerAdapter{})
}

const (
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)
