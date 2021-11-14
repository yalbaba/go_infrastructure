package liveearth_primary

import "time"

// Keywords ...
type Keywords struct {
	// ID 关键词id
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// Keyword 关键词
	Keyword string `json:"keyword"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// UpdateTime 更新时间
	UpdateTime time.Time `json:"update_time"`

	Sort int `json:"sort"`
}

func (*Keywords) TableName() string {
	return "liveearth_primary.keywords"
}
