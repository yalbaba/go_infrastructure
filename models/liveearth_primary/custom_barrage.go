package liveearth_primary

import "time"

// CustomBarrage ...
type CustomBarrage struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// Content 弹幕内容
	Content string `json:"content"`
	// BarrageType 弹幕类型: 10 = 悲观; 20 = 平淡; 30 = 乐观
	BarrageType int `json:"barrage_type"`
	// Operator 操作人
	Operator string `json:"operator"`
	// UpdateTime 更新时间
	UpdateTime time.Time `json:"update_time"`
}

func (*CustomBarrage) TableName() string {
	return "liveearth_primary.custom_barrage"
}
