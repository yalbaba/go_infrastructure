/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/9/14
   Description :
-------------------------------------------------
*/

package liveearth_primary

type LiveRoomShareProduct struct {
	Id                int    `json:"id"`
	LiveRoomId        int    `json:"live_room_id"`
	CouponId          int    `json:"coupon_id"`   // 优惠券id
	CouponName        string `json:"coupon_name"` // 优惠券名称
	CouponUrl         string `json:"coupon_url"`  // 优惠券链接
	CoverUrl          string `json:"cover_url"`
	Title             string `json:"title"`
	StartTime         int    `json:"start_time"`         // 领取开始时间
	EndTime           int    `json:"end_time"`           // 领取结束时间
	CouponType        int    `json:"coupon_type"`        // 优惠券类型; 10=现金券; 20=折扣券
	PreferentialQuota int    `json:"preferential_quota"` // 优惠券额度
	BasePeopleNum     int    `json:"base_people_num"`    // 基础人数
}

func (*LiveRoomShareProduct) TableName() string {
	return "liveearth_primary.live_room_share_product"
}
