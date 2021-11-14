package liveearth_primary

import "time"

type LiveRoomCollectionRelationship struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	CollectionId int       `json:"collection_id" `  // 集合id
	LiveRoomId   int64     `json:"live_room_id" `   // 直播id
	LiveRoomName string    `json:"live_room_name" ` // 直播名
	SortId       int8      `json:"sort_id" `        // 排序id
	CreateTime   time.Time `json:"create_time" `    // 创建时间
}

func (m *LiveRoomCollectionRelationship) TableName() string {
	return "liveearth_primary.live_room_collection_relationship"
}
