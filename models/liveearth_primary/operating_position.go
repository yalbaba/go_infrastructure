package liveearth_primary

import "time"

// OperatingPosition ...
type OperatingPosition struct {
	Id int64 `json:"id"`
	// Title 内容标题
	Title string `json:"title"`
	// ContentId 内容id
	ContentId int64 `json:"content_id"`
	// ContentType 内容类型: 8 = 视频; 1 = 直播; 32 = 话题; 4 = 发现
	ContentType int `json:"content_type"`
	// ColumnType 栏目类型: 推荐 = 99, 直播 = 1, 视频 = 8, 发现 = 4
	ColumnType int `json:"column_type"`
	// State 状态: 10 = 不生效; 20 = 生效
	State int `json:"state"`
	// ExposureNum 曝光次数
	ExposureNum int `json:"exposure_num"`
	// StartTime 有效期开始时间
	StartTime time.Time `json:"start_time"`
	// EndTime 有效期结束时间
	EndTime time.Time `json:"end_time"`
	// UpdateTime 编辑时间
	UpdateTime time.Time `json:"update_time"`
	// Operator 操作人
	Operator string `json:"operator"`
	// SortId 排序, 小的排前面
	SortId int `json:"sort_id"`
}

func (*OperatingPosition) TableName() string {
	return "liveearth_primary.operating_position"
}
