package liveearth_primary

import (
	"time"
)

type LiveRoomSpecialSubjectRelationship struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveRoomId       int       `json:"live_room_id" xorm:"not null comment('直播间id') INT(10)"`
	SpecialSubjectId int       `json:"special_subject_id" xorm:"not null comment('专题id') INT(10)"`
	CreateTime       time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*LiveRoomSpecialSubjectRelationship) TableName() string {
	return "liveearth_primary.live_room_special_subject_relationship"
}
