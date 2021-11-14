package guide

import "time"

type ActivityUserCoupon struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId     string    `json:"user_id" `     // 用户的id
	ActivityId int       `json:"activity_id" ` // 活动id
	PrizeId    int       `json:"prize_id" `    // 抽中的奖品id
	CreateTime time.Time `json:"create_time" ` // 创建时间
}

func (m *ActivityUserCoupon) TableName() string {
	return "guide.activity_user_coupon"
}
