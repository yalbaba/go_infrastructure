package liveearth_primary

import (
	"time"
)

type UserMessage struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Receiver string `json:"receiver" xorm:"not null comment('消息接收者ID') VARCHAR(32)"`
	Sender   string `json:"sender" xorm:"not null comment('消息发送者ID， 系统消息为0') VARCHAR(32)"`
	Content  string `json:"content" xorm:"default '' comment('消息内容') VARCHAR(255)"`
	Reason   int    `json:"reason" xorm:"not null comment('被通知的原因
10: 被关注
20: ugc内容(故事)被点赞
30: ugc内容(故事)被收藏
40: ugc内容(故事)被评论
50: 评论被回复
60: 评论被点赞
70: 审核未通过
80: 话题上热门
') TINYINT(4)"`
	MessageType    int       `json:"message_type" xorm:"not null comment('消息类型 10:系统消息(站内信) 20:用户互动消息') TINYINT(4)"`
	MessageStatus  int       `json:"message_status" xorm:"not null default 10 comment('消息状态 0:删除 10:未读 20:已读') TINYINT(4)"`
	CreateTime     time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime     time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	ArticleId      string    `json:"article_id" xorm:"comment('文章、故事、主播分享的内容等ID') VARCHAR(32)"`
	ArticleType    int       `json:"article_type" xorm:"not null default 0 comment('文章类型 10 资讯 20 故事 30 主播发言分享') TINYINT(4)"`
	Url            string    `json:"url"`
	TopicName      string    `json:"topic_name"`
	Cover          string    `json:"cover"`
	AdminMessageId int       `json:"admin_message_id"`
	IsClick        int       `json:"is_click"`
	IsExposure     int       `json:"is_exposure"`
}

func (*UserMessage) TableName() string {
	return "liveearth_primary.user_message"
}
