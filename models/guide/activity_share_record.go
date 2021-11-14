package guide

import "time"

type ActivityShareRecord struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ActivityId     int       `json:"activity_id" `      // 活动id
	ShareUserId    string    `json:"share_user_id" `    // 分享活动的用户id
	RegisterUserId string    `json:"register_user_id" ` // 通过活动注册的用户id
	CreateTime     time.Time `json:"create_time" `      // 创建时间
}

func (m *ActivityShareRecord) TableName() string {
	return "guide.activity_share_record"
}
