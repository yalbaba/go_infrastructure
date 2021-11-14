package liveearth_primary

import "time"

// LiveSourceTag ...
type LiveSourceTag struct {
	Id uint64 `json:"id" xorm:"not null pk autoincr INT(10)"`
	// LiveSourceId 直播源id
	LiveSourceId uint64 `json:"live_source_id"`
	// TagID 标签id
	TagId      uint64    `json:"tag_id"`
	CreateTime time.Time `json:"create_time"`
	TagLevel   int       `json:"tag_level"`
}

func (*LiveSourceTag) TableName() string {
	return "liveearth_primary.live_source_tag"
}
