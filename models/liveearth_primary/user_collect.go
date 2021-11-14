package liveearth_primary

import (
	"time"
)

type UserCollect struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId     string    `json:"user_id" xorm:"not null unique(user_collect_user_id_type_related_id_uindex) VARCHAR(32)"`
	Type       int       `json:"type" xorm:"not null comment('内容类型  10 = 直播; 20 = 足迹; 30 = 百科') unique(user_collect_user_id_type_related_id_uindex) TINYINT(4)"`
	RelatedId  int64     `json:"related_id" xorm:"not null default 0 comment('关联id') unique(user_collect_user_id_type_related_id_uindex) BIGINT(20)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*UserCollect) TableName() string {
	return "liveearth_primary.user_collect"
}
