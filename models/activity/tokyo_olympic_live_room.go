package activity

import "time"

type TokyoOlympicLiveRoom struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveRoomId   int64     `json:"live_room_id" gorm:"column:live_room_id"`     // 直播间id
	LiveRoomName string    `json:"live_room_name" gorm:"column:live_room_name"` // 直播间名称
	Introduction string    `json:"introduction" gorm:"column:introduction"`     // 直播间简介
	State        int8      `json:"state"`
	Sort         int       `json:"sort" gorm:"column:sort"`               // 排序
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
}

func (m *TokyoOlympicLiveRoom) TableName() string {
	return "activity.tokyo_olympic_live_room"
}
