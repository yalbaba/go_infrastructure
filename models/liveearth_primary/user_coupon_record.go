package liveearth_primary

import "time"

type UserCouponRecord struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr BIGINT(20)"`
	PhoneNumber   string    `json:"phone_number"`    //领取用户的手机号
	CouponId      int64     `json:"coupon_id"`       //优惠券id
	Title         string    `json:"title"`           //优惠券名称
	TakeTime      time.Time `json:"take_time"`       //领取时间
	CouponGroupId int64     `json:"coupon_group_id"` //优惠券活动id
	UserId        string    `json:"user_id"`         //用户id
}

func (*UserCouponRecord) TableName() string {
	return "liveearth_primary.user_coupon_record"
}
