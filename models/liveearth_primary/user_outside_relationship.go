package liveearth_primary

import "time"

type UserOutsideRelationship struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr INT(10)"` // 自增id
	UserId        string    `json:"user_id" `                               // 用户id
	PhoneNumber   string    `json:"phone_number" `                          // 用户手机号码
	WechatName    string    `json:"wechat_name" `                           // 微信名
	WechatAvatar  string    `json:"wechat_avatar" `                         // 微信头像
	WechatSex     int8      `json:"wechat_sex" `                            // 微信性别 1 为男性，2 为女性,3:不明
	WechatOpenId  string    `json:"wechat_open_id" `                        // 微信openId
	WechatUnionId string    `json:"wechat_union_id" `                       // 微信unionId
	AppleUserId   string    `json:"apple_user_id"`                          // 苹果获取user_id
	PlatformType  int8      `json:"platform_type" `                         // 接入第三方类型(10:微信app，20:小程序，30苹果登录)
	LoginTime     time.Time `json:"login_time" `                            // 登录时间
	CreateTime    time.Time `json:"create_time" `                           // 创建时间
}

func (m *UserOutsideRelationship) TableName() string {
	return "liveearth_primary.user_outside_relationship"
}
