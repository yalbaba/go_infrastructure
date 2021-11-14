package liveearth_primary

import (
	"time"
)

type TopicOnList struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId      string    `json:"user_id" xorm:"not null comment('用户id') unique(earth_topic_on_list_user_id_story_id_topic_id_uindex) VARCHAR(32)"`
	FootprintId int       `json:"footprint_id" xorm:"not null comment('足迹id') unique(earth_topic_on_list_user_id_story_id_topic_id_uindex) INT(10)"`
	TopicId     int       `json:"topic_id" xorm:"not null comment('话题id') unique(earth_topic_on_list_user_id_story_id_topic_id_uindex) INT(10)"`
	CreateTime  time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}

func (*TopicOnList) TableName() string {
	return "liveearth_primary.topic_on_list"
}
