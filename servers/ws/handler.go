package ws

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yalbaba/go_infrastructure/pkg/iris"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type wsHandler struct {
	// The websocket connection.
	conn      *websocket.Conn
	closeChan chan struct{}
	once      sync.Once
	// Buffered channel of outbound messages.
	send     chan []byte
	jwtToken string
}

func newWSHandler(conn *websocket.Conn) *wsHandler {
	return &wsHandler{
		conn:      conn,
		closeChan: make(chan struct{}),
		send:      make(chan []byte, 256),
	}
}

//ws handler
type Handler func(ctx iris.Context, message []byte) interface{}

//readPump 循环从读取客户端传入数据
func (c *wsHandler) readPump(hf Handler, ctx iris.Context) {
	defer func() {
		c.close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		// Read
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			ctx.Error("ReadMessage", err)
			break
		}

		fmt.Printf("ws get message: %s\n", msg)

		// 执行业务函数
		r := hf(ctx, msg)

		// Write
		var res []byte
		switch r.(type) {
		case string:
			res = []byte(r.(string))
		case []byte:
			res = r.([]byte)
		default:

		}

		err = c.conn.WriteMessage(websocket.TextMessage, res)
		if err != nil {
			ctx.Error("WriteMessage", err)
			break
		}

	}
}

func (c *wsHandler) writePump() {
	ticker := time.NewTicker(time.Second * 60)
	defer func() {
		ticker.Stop()
		c.close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.close()
				return
			}
			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				c.close()
				break
			}
			w.Write(message)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				c.close()
				break
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				c.close()
				break
			}
			fmt.Println("ping")
		case <-c.closeChan:
			break
		}
	}
}

func (c *wsHandler) recvNotify(ctx iris.Context) func(buff []byte) error {
	return func(buff []byte) error {
		c.send <- buff
		return nil
	}
}

func (c *wsHandler) close() {
	c.once.Do(func() {
		close(c.closeChan)
		close(c.send)
	})
}

func getUUID(c iris.Context) string {

	if v, ok := c.Values().Get("__parrot_sid_").(string); ok {
		return v
	}

	ck, err := c.Request().Cookie("parrot_sid")
	if err != nil || ck == nil || ck.Value == "" {
		return c.GetSessionID()
	}

	return ck.Value
}
