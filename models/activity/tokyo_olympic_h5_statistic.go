package activity

import "time"

type TokyoOlympicH5Statistic struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId     string    `json:"user_id" gorm:"column:user_id"`         // 用户id
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
}

func (m *TokyoOlympicH5Statistic) TableName() string {
	return "activity.tokyo_olympic_h5_statistic"
}
