package liveearth_primary

import "time"

//小程序推送消息记录表
type PullNewActivityNotifyRecord struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr INT(20)"` //主键
	UserId     string    `json:"user_id"`                                //用户id
	UserName   string    `json:"user_name"`                              //用户名
	Page       string    `json:"page"`                                   //点击模板卡片后的跳转页面
	TemplateId string    `json:"template_id"`                            //订阅模板id
	CreateTime time.Time `json:"create_time"`                            //创建时间
}

func (*PullNewActivityNotifyRecord) TableName() string {
	return "liveearth_primary.pull_new_activity_notify_record"
}
