package liveearth_primary

import "time"

type LiveRoomRecommendManage struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveRoomId  int       `json:"live_room_id"`
	State       int8      `json:"state" gorm:"column:state"`               // 10 关闭 20开启
	Sort        int       `json:"sort" gorm:"column:sort"`                 // 排序
	ContentType int8      `json:"content_type" gorm:"column:content_type"` // 内容类型  10 直播  20 视频
	ContentId   int       `json:"content_id" gorm:"column:content_id"`     // 内容对应的id
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *LiveRoomRecommendManage) TableName() string {
	return "liveearth_primary.live_room_recommend_manage"
}
