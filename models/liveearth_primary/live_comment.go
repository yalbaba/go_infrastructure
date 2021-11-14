package liveearth_primary

import (
	"time"
)

type LiveComment struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveRoomId      int    `json:"live_room_id" xorm:"not null comment('聊天室id') index INT(10)"`
	LiveSourceId    int    `json:"live_source_id" xorm:"not null INT(10)"`
	UserId          string `json:"user_id" xorm:"not null default '' comment('发言人用户id') VARCHAR(32)"`
	UserName        string `json:"user_name" xorm:"not null default '' comment('发言人用户名') VARCHAR(32)"`
	UserPhoneNumber string `json:"user_phone_number" xorm:"comment('发言人电话') VARCHAR(16)"`
	Content         string `json:"content" xorm:"not null default '' comment('发言内容') VARCHAR(128)"`
	CoverUrl        string `json:"cover_url" xorm:"comment('视频的封面图') VARCHAR(1024)"`
	MediaUrl        string `json:"media_url" xorm:"not null default '' comment('媒体url
') VARCHAR(1024)"`
	MediaType   int       `json:"media_type" xorm:"not null default 10 comment('媒体类型  10 = 没有媒体资源;20 = 图片; 30 = 视频') TINYINT(4)"`
	CommentType int       `json:"comment_type" xorm:"not null default 10 comment(' 发言类型 10 = 精彩评论') TINYINT(4)"`
	State       int       `json:"state" xorm:"not null default 10 comment(' 10 = 未推荐; 20 = 推荐') TINYINT(4)"`
	CommentTime time.Time `json:"comment_time" xorm:"not null default CURRENT_TIMESTAMP comment('发言时间') index TIMESTAMP"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime  time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Operator    string    `json:"operator" xorm:"not null default '' comment('操作人') VARCHAR(16)"`
}

func (*LiveComment) TableName() string {
	return "liveearth_primary.live_comment"
}
