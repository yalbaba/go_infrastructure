package activity

import "time"

type TokyoOlympicBase struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	TopicName  string    `json:"topic_name" gorm:"column:topic_name"`   // 直播主题名称
	Url        string    `json:"url" gorm:"column:url"`                 // 海报地址
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
}

func (m *TokyoOlympicBase) TableName() string {
	return "activity.tokyo_olympic_base"
}
