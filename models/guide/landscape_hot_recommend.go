package guide

import "time"

type LandscapeHotRecommend struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	LandscapeViewId int       `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区id
	SortId          int       `json:"sort_id"`
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
	Operator        string    `json:"operator" gorm:"column:operator"` // 操作人名称
}

func (m *LandscapeHotRecommend) TableName() string {
	return "guide.landscape_hot_recommend"
}
