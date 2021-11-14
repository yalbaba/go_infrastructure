package liveearth_primary

import "time"

type BirdsViewDeviceRelationship struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	BirdsViewDeviceId int       `json:"birds_view_device_id" gorm:"column:birds_view_device_id"` // 鸟瞰设备id
	LiveSourceId      int64     `json:"live_source_id" gorm:"column:live_source_id"`             // 直播源id
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`
}

func (m *BirdsViewDeviceRelationship) TableName() string {
	return "liveearth_primary.birds_view_device_relationship"
}
