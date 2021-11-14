package liveearth_primary

import "time"

// HomePageColumn ...
type HomePageColumn struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(11)"`
	// ColumnName 首页栏目名称
	ColumnName string `json:"column_name"`
	// ColumnType 所属模块 0 空 10 关注 20 推荐 30 直播 40 视频 50 发现 60 卫星TV 70 话题 80 足迹
	ColumnType int `json:"column_type"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// UpdateTime 更新时间
	UpdateTime time.Time `json:"update_time"`
	// Operator 最近一次的操作人
	Operator string `json:"operator"`
	// Sort 栏目位置排序
	Sort int `json:"sort"`

	State int `json:"state"`

	IsRecommend int `json:"is_recommend"`

	Deleted int `json:"deleted"`
}

func (*HomePageColumn) TableName() string {
	return "liveearth_primary.home_page_column"
}
