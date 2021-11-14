package guide

import "time"

type LandscapeLiveRecommend struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	LiveRoomId      int64     `json:"live_room_id" gorm:"column:live_room_id"`
	LandscapeViewId int       `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区id
	SortId          int       `json:"sort_id"`
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
	Operator        string    `json:"operator" gorm:"column:operator"` // 操作人名称
}

func (m *LandscapeLiveRecommend) TableName() string {
	return "guide.landscape_live_recommend"
}
