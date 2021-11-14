package liveearth_primary

import (
	"time"
)

type AdminMessagePush struct {
	Id              int `json:"id" xorm:"not null pk autoincr comment('pk') INT(10)"`
	TemplateId      int `json:"template_id" xorm:"not null default 0 comment('使用的推送模板ID') INT(10)"`
	MessageLocation int `json:"message_location" xorm:"not null default 10 comment('消息位置 10 app外(通知栏) 20 app内') TINYINT(4)"`
	MessageStatus   int `json:"message_status" xorm:"not null default 10 comment('消息状态 0 删除, 10 未删除') TINYINT(4)"`
	//MessageType      int       `json:"message_type" xorm:"not null default 10 comment('消息类型 10 无(打开app), 20 url链接, 30 资讯, 40 Ugc, 50 直播') TINYINT(4)"`
	Content string `json:"content" xorm:"not null default '' comment('内容uri: 资讯、ugc、直播间ID、url链接等 ') VARCHAR(1024)"`
	//AcceptArea       string    `json:"accept_area" xorm:"not null comment('可接收消息范围') JSON"`
	//AcceptDeviceType int       `json:"accept_device_type" xorm:"not null default 0 comment('消息接收设备类型 10 android, 20 ios, 30 全部') TINYINT(4)"`
	PushTitle       string    `json:"push_title" xorm:"not null comment('推送的标题') VARCHAR(64)"`
	PushBody        string    `json:"push_body" xorm:"not null comment('推送的body') VARCHAR(255)"`
	PushTime        time.Time `json:"push_time" xorm:"not null default CURRENT_TIMESTAMP comment('推送时间') TIMESTAMP"`
	PushDeviceCount int       `json:"push_device_count" xorm:"not null default 0 comment('推送的设备数') INT(10)"`
	PushSend        int       `json:"push_send" xorm:"not null default 0 comment('推送发出数') INT(10)"`
	PushArrive      int       `json:"push_arrive" xorm:"not null default 0 comment('推送抵达数') INT(10)"`
	PushClick       int       `json:"push_click" xorm:"not null default 0 comment('推送点击数') INT(10)"`
	PushClear       int       `json:"push_clear" xorm:"not null default 0 comment('推送清除数') INT(10)"`
	//PushStatus      int       `json:"push_status" xorm:"not null comment('推送状态 10 未推送 20 已推送 30 禁用') TINYINT(4)"`
	//PushParams string    `json:"push_params" xorm:"not null default '' comment('推送的额外参数') VARCHAR(2048)"`
	Creator       string    `json:"creator" xorm:"not null comment('推送的创建者') VARCHAR(32)"`
	Editor        string    `json:"editor" xorm:"not null default '' comment('推送的编辑者') VARCHAR(32)"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime    time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	CoverUrl      string    `json:"cover_url"`
	TargetUser    string    `json:"target_user"`
	AppPushAction int       `json:"app_push_action"`
}

func (*AdminMessagePush) TableName() string {
	return "liveearth_primary.admin_message_push"
}
