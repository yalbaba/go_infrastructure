package liveearth_primary

import (
	"time"
)

type LiveAnchorComment struct {
	Id             int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveRoomId     int    `json:"live_room_id" xorm:"not null comment('聊天室id') index INT(10)"`
	LiveCommentId  int    `json:"live_comment_id" xorm:"not null default 0 comment('实时评论表的id') INT(10)"`
	UserId         string `json:"user_id" xorm:"not null comment('直播间所有者id') VARCHAR(32)"`
	UserName       string `json:"user_name" xorm:"not null comment('直播间所有者用户名') VARCHAR(32)"`
	AnchorId       string `json:"anchor_id" xorm:"not null comment('主持人id') VARCHAR(32)"`
	AnchorName     string `json:"anchor_name" xorm:"not null comment('主持人用户名') VARCHAR(32)"`
	Type           int    `json:"type" xorm:"not null comment(' 10 = 主持人推荐; 20 = 主持人自主发言') TINYINT(4)"`
	Content        string `json:"content" xorm:"comment('内容') VARCHAR(255)"`
	Recommendation string `json:"recommendation" xorm:"not null default '' comment('推荐语
') VARCHAR(255)"`
	CoverUrl string `json:"cover_url" xorm:"comment('视频封面图片') VARCHAR(1024)"`
	MediaUrl string `json:"media_url" xorm:"comment('媒体url
') VARCHAR(1024)"`
	MediaType  int       `json:"media_type" xorm:"not null default 10 comment('媒体类型 10 = 没有; 20 = 图片; 30 = 视频') TINYINT(4)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP index TIMESTAMP"`
	UpdateTime time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*LiveAnchorComment) TableName() string {
	return "liveearth_primary.live_anchor_comment"
}
