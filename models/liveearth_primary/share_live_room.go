package liveearth_primary

import "time"

type ShareLiveRoom struct {
	Id         int32     `json:"id"`
	SignalName string    `json:"signal_name"`
	UserId     string    `json:"user_id"`
	DeviceId   string    `json:"device_id"`
	LiveRoomId int64     `json:"live_room_id"`
	ShareTime  time.Time `json:"share_time"`
	UserIp     string    `json:"user_ip"`
}

func (*ShareLiveRoom) TableName() string {
	return "liveearth_primary.share_live_room"
}
