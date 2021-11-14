package liveearth_primary

import "time"

// ColumnTags ...
type ColumnTags struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr INT(10)"`
	ColumnsId  int64     `json:"columns_id" xorm:"columns_id"`
	TagId      int64     `json:"tag_id" xorm:"tag_id"`
	CreateTime time.Time `json:"create_time" xorm:"create_time"`
	TagLevel   int       `json:"tag_level"`
}

func (*ColumnTags) TableName() string {
	return "liveearth_primary.column_tags"
}
