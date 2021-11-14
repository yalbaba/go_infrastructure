package liveearth_primary

type BaseWxTemplateInfo struct {
	Id         int64  `json:"id" xorm:"not null pk autoincr INT(20)"`
	TemplateId string `json:"template_id"` //微信通知消息模板id
	UseTimes   int    `json:"use_times"`   //模板剩余使用次数
}
