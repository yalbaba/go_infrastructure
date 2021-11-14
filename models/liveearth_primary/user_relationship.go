package liveearth_primary

import (
	"time"
)

type UserRelationship struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId      string    `json:"user_id" xorm:"not null comment('用户id') VARCHAR(32)"`
	FollowerUid string    `json:"follower_uid" xorm:"not null comment('粉丝id') VARCHAR(32)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*UserRelationship) TableName() string {
	return "liveearth_primary.user_relationship"
}
