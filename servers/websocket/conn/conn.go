package conn

import (
	"github.com/gorilla/websocket"
)

type IConn interface {
	GetID() string
	GetGroupId() int64
	GetDeviceId() string
	GetConn() *websocket.Conn
	GetPlatform() int
	Close() error
	Send(event string, input interface{}) error
}

var (
	Android int = 20 //安卓
	Ios     int = 30 //苹果
	MiniApp int = 50 //小程序
)

//websocket的连接对象
type Conn struct {
	ID       string
	GroupId  int64
	Conn     *websocket.Conn
	DeviceId string //小程序是UserId、app是设备id
	Platform int
}

func NewConn(id string, groupId int64, deviceId string, platform int, conn *websocket.Conn) *Conn {
	return &Conn{
		ID:       id,
		GroupId:  groupId,
		Conn:     conn,
		Platform: platform,
		DeviceId: deviceId,
	}
}

func (c *Conn) GetID() string {
	return c.ID
}

func (c *Conn) GetConn() *websocket.Conn {
	return c.Conn
}

func (c *Conn) Send(event string, result interface{}) error {
	return c.Conn.WriteJSON(WriteResponse(event, result))
}

func (c *Conn) GetDeviceId() string {
	return c.DeviceId
}

func (c *Conn) GetGroupId() int64 {
	return c.GroupId
}

func (c *Conn) GetPlatform() int {
	return c.Platform
}

func (c *Conn) Close() error {
	return c.Conn.Close()
}
