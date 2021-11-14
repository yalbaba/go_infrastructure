package guide

import "time"

type ActivityStatistics struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Pv       int       `json:"pv" `        // 每日的pv
	Uv       int       `json:"uv" `        // 每日的uv
	PlatForm int8      `json:"plat_form" ` // 平台：10，app端，20，h5端
	Date     time.Time `json:"date" `      // 日期
}

func (m *ActivityStatistics) TableName() string {
	return "guide.activity_statistics"
}
