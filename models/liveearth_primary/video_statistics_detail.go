package liveearth_primary

import "time"

type VideoStatisticsDetail struct {
	Id       int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	VideoId  int64     `json:"video_id"`
	VideoUv  int       `json:"video_uv"`
	VideoPv  int       `json:"video_pv"`
	CreateAt time.Time `json:"create_at"`
}

func (*VideoStatisticsDetail) TableName() string {
	return "liveearth_primary.video_statistics_detail"
}
