package liveearth_primary

import "time"

//小程序推送消息基本信息表
type PullNewActivityNotifyInfo struct {
	Id           int64     `json:"id" xorm:"not null pk autoincr INT(20)"` //主键
	UserId       string    `json:"user_id"`                                //用户id
	WechatOpenId string    `json:"wechat_open_id"`                         //OpenId
	TemplateId   string    `json:"template_id"`                            //订阅模板id
	CreateTime   time.Time `json:"create_time"`                            //创建时间
}

func (*PullNewActivityNotifyInfo) TableName() string {
	return "liveearth_primary.pull_new_activity_notify_info"
}
