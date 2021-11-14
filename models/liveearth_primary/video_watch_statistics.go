package liveearth_primary

import "time"

type VideoWatchStatistics struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	VideoId    int64     `json:"video_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	UserId     string    `json:"user_id"`
	WatchTime  float64   `json:"watch_time"`
	CreateTime time.Time `json:"create_time"`
}

func (*VideoWatchStatistics) TableName() string {
	return "liveearth_primary.video_watch_statistics"
}
