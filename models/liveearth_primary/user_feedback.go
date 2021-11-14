package liveearth_primary

import (
	"time"
)

type UserFeedback struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId     string    `json:"user_id" xorm:"not null VARCHAR(32)"`
	Contact    string    `json:"contact" xorm:"not null comment('联系人') VARCHAR(8)"`
	Content    string    `json:"content" xorm:"not null comment('意见') VARCHAR(255)"`
	ContactWay string    `json:"contact_way" xorm:"not null comment('联系方式: 电话等...') VARCHAR(32)"`
	Image      string    `json:"image" xorm:"comment('故障截图数组(json)') VARCHAR(1024)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*UserFeedback) TableName() string {
	return "liveearth_primary.user_feedback"
}
