package liveearth_primary

import "time"

// VideoTag ...
type VideoTag struct {
	// ID 数据id
	Id int `json:"id"`
	// VideoID 视频id
	VideoId int64 `json:"video_id"`
	// TagID 标签id
	TagId int `json:"tag_id"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`

	TagLevel int `json:"tag_level"`
}

func (*VideoTag) TableName() string {
	return "liveearth_primary.video_tag"
}
