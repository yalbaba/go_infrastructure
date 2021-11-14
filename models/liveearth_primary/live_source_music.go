package liveearth_primary

import "time"

type LiveSourceMusic struct {
	Id         int       `json:"id"  xorm:"not null pk autoincr INT(10)"`
	Name       string    `json:"name" xorm:"name"`
	MusicType  int       `json:"music_type" xorm:"music_type"`
	Url        string    `json:"url" xorm:"url"`
	CoverUrl   string    `json:"cover_url" xorm:"cover_url" `
	CreateTime time.Time `json:"create_time" xorm:"create_time"`
}

func (*LiveSourceMusic) TableName() string {
	return "liveearth_primary.live_source_music"
}
