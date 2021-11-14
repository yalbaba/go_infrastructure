package liveearth_primary

import "time"

type TopicLiveRoomRelationship struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	TopicId    int64     `json:"topic_id"`
	LiveRoomId int64     `json:"live_room_id"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `json:"create_time"`
}

func (*TopicLiveRoomRelationship) TableName() string {
	return "liveearth_primary.topic_live_room_relationship"
}
