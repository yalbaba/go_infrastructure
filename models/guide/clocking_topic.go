package guide

import "time"

type ClockingTopic struct {
	Id         int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"`
	ClockingId int       `json:"clocking_id" gorm:"column:clocking_id"`
	TopicId    int       `json:"topic_id" gorm:"column:topic_id"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

func (m *ClockingTopic) TableName() string {
	return "guide.clocking_topic"
}
