package liveearth_primary

import "time"

type ShareLiveRoomDetail struct {
	Id int32 `json:"id"`
	// UUID 每次链接的唯一标识符
	SignalName string `json:"uuid"`
	// DeviceID 设备id
	DeviceId string `json:"device_id"`
	// UserID 用户id
	UserId    string    `json:"user_id"`
	ClickTime time.Time `json:"click_time"`
	UserIp    string    `json:"user_ip"`
	ShareType int       `json:"share_type"` //分享类型
}

func (*ShareLiveRoomDetail) TableName() string {
	return "liveearth_primary.share_live_room_detail"
}
