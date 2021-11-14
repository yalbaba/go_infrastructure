package liveearth_primary

import (
	"time"
)

type User struct {
	Id            int64     `json:"id" xorm:"pk autoincr comment('自增id') BIGINT(20)"`
	UserId        string    `json:"user_id" xorm:"not null comment('用户id') index VARCHAR(32)"`
	PhoneNumber   string    `json:"phone_number" xorm:"not null comment('用户手机号码') index VARCHAR(11)"`
	UserName      string    `json:"user_name" xorm:"comment('用户名') VARCHAR(32)"`
	WechatName    string    `json:"wechat_name" xorm:"comment('微信名') VARCHAR(32)"`
	WechatOpenId  string    `json:"wechat_open_id" xorm:"comment('微信openId') index VARCHAR(64)"`
	UnionId       string    `json:"union_id"`
	WechatAvatar  string    `json:"wechat_avatar" xorm:"default '' comment('微信头像') VARCHAR(1024)"`
	WechatSex     int       `json:"wechat_sex" xorm:"comment('微信性别 1 为男性，2 为女性,3:不明') TINYINT(4)"`
	Avatar        string    `json:"avatar" xorm:"comment('用户头像') VARCHAR(1024)"`
	UserDesc      string    `json:"user_desc" xorm:"comment('个人简介') VARCHAR(64)"`
	Gender        int       `json:"gender" xorm:"comment('性别 1 为男性，2 为女性') TINYINT(4)"`
	LoginPlatform int       `json:"login_platform" xorm:"comment('登录平台 50 小程序，60 h5') TINYINT(4)"`
	Birthday      time.Time `json:"birthday" xorm:"comment('出生年月') TIMESTAMP"`
	UserState     int       `json:"user_state" xorm:"comment('用户状态：10: 正常  20: 锁定  30: 注销中') TINYINT(4)"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime    time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	LogoffTime    time.Time `json:"logoff_time"`
	BeLikeNum     int       `json:"be_like_num"`
}

func (*User) TableName() string {
	return "liveearth_primary.user"
}
