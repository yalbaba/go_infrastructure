package liveearth_primary

import "time"

type UserLogoff struct {
	Id              int64     `json:"id" xorm:"pk autoincr comment('自增id') BIGINT(20)"`
	UserId          string    `json:"user_id" xorm:"not null comment('用户id') index VARCHAR(32)"`
	PhoneNumber     string    `json:"phone_number" xorm:"not null comment('用户手机号码') index VARCHAR(11)"`
	UserName        string    `json:"user_name" xorm:"comment('用户名') VARCHAR(32)"`
	IsUgc           int       `json:"is_ugc" xorm:"not null default 0 comment('是否 ugc 用户  10 = 否; 20 = 是') TINYINT(4)"`
	WechatName      string    `json:"wechat_name" xorm:"comment('微信名') VARCHAR(32)"`
	Gender          int       `json:"gender" xorm:"comment('性别 1 为男性，2 为女性') TINYINT(4)"`
	City            string    `json:"city" xorm:"not null default '' comment('城市') VARCHAR(32)"`
	AppInstallTime  time.Time `json:"app_install_time" xorm:"not null default CURRENT_TIMESTAMP comment('app安装时间') TIMESTAMP"`
	UserCreateTime  time.Time `json:"user_create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	LiveAuth        int       `json:"live_auth" xorm:"not null default 0 comment('直播作者认证') TINYINT(4)"`
	StoryAuthorAuth int       `json:"story_author_auth" xorm:"not null default 0 comment('故事作者认证') TINYINT(4)"`
	InformationAuth int       `json:"information_auth" xorm:"not null default 0 comment('资讯作者认证') TINYINT(4)"`
	CreateTime      time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}

func (*UserLogoff) TableName() string {
	return "liveearth_primary.user_logoff"
}
