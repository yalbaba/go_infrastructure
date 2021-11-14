package liveearth_primary

type UserAuth struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId          string `json:"user_id" xorm:"not null unique VARCHAR(32)"`
	IsUgc           int    `json:"is_ugc" xorm:"not null default 0 comment('是否 ugc 用户  10 = 否; 20 = 是') TINYINT(4)"`
	StoryAuthorAuth int    `json:"story_author_auth" xorm:"not null default 0 comment('故事作者认证') TINYINT(4)"`
	InformationAuth int    `json:"information_auth" xorm:"not null default 0 comment('资讯作者认证') TINYINT(4)"`
	LiveAuth        int    `json:"live_auth" xorm:"not null default 0 comment('直播作者认证') TINYINT(4)"`
	AdAuth          int    `json:"ad_auth" xorm:"not null default 0 comment('广告主身份认证') TINYINT(4)"`
	SellAuth        int    `json:"sell_auth" xorm:"not null default 0 comment('店主身份认证') TINYINT(4)"`
}

func (*UserAuth) TableName() string {
	return "liveearth_primary.user_auth"
}
