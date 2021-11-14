package guide

import "time"

type LandscapeVrRecommend struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	VrUrl           string    `json:"vr_url" gorm:"column:vr_url"`                       // vr资源
	LandscapeViewId int       `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区id
	SortId          int       `json:"sort_id"`
	VrVideoId       int64     `json:"vr_video_id"`
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
	Operator        string    `json:"operator" gorm:"column:operator"` // 操作人名称
}

func (m *LandscapeVrRecommend) TableName() string {
	return "guide.landscape_vr_recommend"
}
