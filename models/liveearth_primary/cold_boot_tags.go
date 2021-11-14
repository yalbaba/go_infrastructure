package liveearth_primary

import "time"

// ColdBootTags ...
type ColdBootTags struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(11)"`
	// TagID 标签id
	TagId int `json:"tag_id"`
	// CoverURL 标签图片
	CoverUrl string `json:"cover_url"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// UpdateTime 更新时间
	UpdateTime time.Time `json:"update_time"`
	// Operator 编辑人
	Operator string `json:"operator"`
	// State 是否启用 10 关闭 20开启
	State int `json:"state"`
	// Sort 排序
	Sort int `json:"sort"`
}

func (*ColdBootTags) TableName() string {
	return "liveearth_primary.cold_boot_tags"
}
