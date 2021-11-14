package liveearth_primary

import "time"

type ShareVideo struct {
	Id         int32     `json:"id"`
	SignalName string    `json:"signal_name"`
	UserId     string    `json:"user_id"`
	DeviceId   string    `json:"device_id"`
	VideoId    int64     `json:"video_id"`
	ShareTime  time.Time `json:"share_time"`
	UserIp     string    `json:"user_ip"`
}

func (*ShareVideo) TableName() string {
	return "liveearth_primary.share_video"
}
