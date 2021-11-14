package guide

import "time"

//活动奖品表
type ActivityPrize struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	CouponGroupId int64     `json:"coupon_group_id"` // 有赞优惠券的活动id
	PrizeAmount   int       `json:"prize_amount" `   // 奖品金额（单位：分）
	CreateTime    time.Time `json:"create_time" `    // 创建时间
}

func (m *ActivityPrize) TableName() string {
	return "guide.activity_prize"
}
