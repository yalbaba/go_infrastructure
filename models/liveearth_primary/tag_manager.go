package liveearth_primary

import (
	"time"
)

type TagManager struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	TagName    string    `json:"tag_name" xorm:"tag_name"`
	TagLevel   int32     `json:"tag_level" xorm:"tag_level"`
	ParentId   int64     `json:"parent_id" xorm:"parent_id"`
	Desc       string    `json:"description"  xorm:"description"`
	CreateTime time.Time `json:"create_time" xorm:"create_time"`
	UpdateTime time.Time `json:"update_time" xorm:"update_time"`
	Sort       int       `json:"sort"`
}

func (*TagManager) TableName() string {
	return "liveearth_primary.tag"
}
