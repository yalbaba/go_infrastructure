package liveearth_primary

import "time"

type LiveRoomCollection struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	CollectionName string    `json:"collection_name" ` // 集合名
	SubTitle       string    `json:"sub_title"`        // 直播集合副标题
	CoverUrl       string    `json:"cover_url" `       // 集合背景图
	LiveRoomNum    int       `json:"live_room_num" `   // 关联直播数量
	SortId         int8      `json:"sort_id" `         // 排序id
	State          int8      `json:"state" `           // 状态：10上架，20下架
	CreateTime     time.Time `json:"create_time" `     // 创建时间
}

func (m *LiveRoomCollection) TableName() string {
	return "liveearth_primary.live_room_collection"
}
