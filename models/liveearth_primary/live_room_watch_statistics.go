package liveearth_primary

import "time"

type LiveRoomWatchStatistics struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveRoomId int64     `json:"live_room_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	UserId     string    `json:"user_id"`
	WatchTime  float64   `json:"watch_time"`
	CreateTime time.Time `json:"create_time"`
}

func (*LiveRoomWatchStatistics) TableName() string {
	return "liveearth_primary.live_room_watch_statistics"
}
