package liveearth_primary

import "time"

type LiveClassificationRoomRelationship struct {
	Id                   int       `json:"id" gorm:"column:id"`
	LiveClassificationId int       `json:"live_classification_id" gorm:"column:live_classification_id"` // 直播分类id
	LiveRoomId           int       `json:"live_room_id" gorm:"column:live_room_id"`                     // 直播间id
	CreateTime           time.Time `json:"create_time" gorm:"column:create_time"`                       // 创建时间
	UpdateTime           time.Time `json:"update_time" gorm:"column:update_time"`                       // 更新时间
}

func (m *LiveClassificationRoomRelationship) TableName() string {
	return "liveearth_primary.live_classification_room_relationship"
}
