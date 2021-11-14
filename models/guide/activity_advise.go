package guide

import "time"

type ActivityAdvise struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Content    string    `json:"content" `     // 建议内容
	CreateTime time.Time `json:"create_time" ` // 创建时间
}

func (m *ActivityAdvise) TableName() string {
	return "guide.activity_advise"
}
