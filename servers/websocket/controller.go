package websocket

import "liveearth/infrastructure/servers/websocket/conn"

type MessageHandle func(input interface{}) interface{}

type ConnController interface {
	OnConnect(conn conn.IConn) (interface{}, error)
	OnClose(conn conn.IConn) error
}
