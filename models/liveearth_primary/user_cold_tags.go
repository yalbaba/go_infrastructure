package liveearth_primary

import "time"

// UserColdTags ...
type UserColdTags struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// TagID 标签id
	TagId int `json:"tag_id"`
	// DeviceID 设备id
	DeviceId   string    `json:"device_id"`
	CreateTime time.Time `json:"create_time"`
}

func (*UserColdTags) TableName() string {
	return "liveearth_primary.user_cold_tags"
}
