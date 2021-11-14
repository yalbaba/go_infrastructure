package liveearth_primary

import "time"

type LiveRoomFloatingInformation struct {
	Id           int       `json:"id" gorm:"column:id"`
	State        int8      `json:"state" gorm:"column:state"`                 // 开启状态 10 关闭， 20 开启
	LiveRoomId   int64     `json:"live_room_id" gorm:"column:live_room_id"`   // 直播间id
	Content      string    `json:"content" gorm:"column:content"`             // 文字内容
	WaitInterval int       `json:"wait_interval" gorm:"column:wait_interval"` // 间隔时间
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`     // 创建时间
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`     // 更新时间
	Deleted      int       `json:"deleted"  gorm:"column:deleted"`            // 是否删除 10 否， 20 是
}

func (m *LiveRoomFloatingInformation) TableName() string {
	return "liveearth_primary.live_room_floating_information"
}
