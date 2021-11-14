package liveearth_primary

import "time"

// EncyclopediaTagRelationship ...
type EncyclopediaTagRelationship struct {
	// ID pk
	Id int `json:"id" xorm:"not null pk autoincr INT(11)"`
	// EncyclopediaID 百科id
	EncyclopediaId int `json:"encyclopedia_id"`
	// TagID 标签id
	TagId int `json:"tag_id"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// TagLevel 标签级别
	TagLevel int `json:"tag_level"`
}

func (*EncyclopediaTagRelationship) TableName() string {
	return "liveearth_primary.encyclopedia_tag_relationship"
}
