package guide

import "time"

//活动锦鲤报名表
type ActivityLuckyPartake struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ActivityId int       `json:"activity_id" ` // 活动id
	UserId     string    `json:"user_id" `     // 锦鲤抽奖用户id
	IsLucky    int8      `json:"is_lucky" `    // 是否中锦鲤，10：否，20：是
	CreateTime time.Time `json:"create_time" ` // 创建时间
}

func (m *ActivityLuckyPartake) TableName() string {
	return "guide.activity_lucky_partake"
}
