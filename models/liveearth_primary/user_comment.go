package liveearth_primary

import (
	"time"
)

type UserComment struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('pk') INT(10)"`
	Content       string    `json:"content" xorm:"not null comment('评论内容') VARCHAR(255)"`
	CommentUserId string    `json:"comment_user_id" xorm:"not null comment('评论用户ID') VARCHAR(32)"`
	ArticleId     string    `json:"article_id" xorm:"not null comment('文章、故事、主播分享内容ID') VARCHAR(32)"`
	ArticleType   int       `json:"article_type" xorm:"not null comment('文章类型 10 资讯 20 故事 30 主播发言分享') TINYINT(4)"`
	AuditTime     time.Time `json:"audit_time" xorm:"not null default '0000-00-00 00:00:00' comment('人工审核时间') TIMESTAMP"`
	AuditUserId   string    `json:"audit_user_id" xorm:"default '' comment('审核人ID') VARCHAR(32)"`
	AuditUserName string    `json:"audit_user_name" xorm:"default '' comment('审核人名称') VARCHAR(32)"`
	AuthorId      string    `json:"author_id" xorm:"default '' comment('作者ID') VARCHAR(32)"`
	ParentId      int       `json:"parent_id" xorm:"not null default 0 comment('父评论id，一级评论为0') INT(10)"`
	LikeNum       int       `json:"like_num" xorm:"not null default 0 comment('点赞数') INT(10)"`
	ReplyNum      int       `json:"reply_num" xorm:"not null default 0 comment('回复数') INT(10)"`
	CreateTime    time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	CommentTime   time.Time `json:"comment_time" xorm:"not null default CURRENT_TIMESTAMP comment('评论时间') TIMESTAMP"`
	CommentStatus int       `json:"comment_status" xorm:"not null default 10 comment('评论状态 0 删除 10 未审核 20 初审通过 30 复审通过 40 初审未通过 50 复审未通过') TINYINT(4)"`
}

func (*UserComment) TableName() string {
	return "liveearth_primary.user_comment"
}
