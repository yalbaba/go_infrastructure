package activity

import "time"

type TokyoOlympicUserAssistance struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId         string    `json:"user_id" gorm:"column:user_id"`                 // 用户id
	AssistanceSort int       `json:"assistance_sort" gorm:"column:assistance_sort"` // 助力排位
	Province       string    `json:"province" gorm:"column:province"`               // 助力省份
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time"`         // 创建时间
}

func (m *TokyoOlympicUserAssistance) TableName() string {
	return "activity.tokyo_olympic_user_assistance"
}
