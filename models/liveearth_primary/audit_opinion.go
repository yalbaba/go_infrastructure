/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/24
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type AuditOpinion struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Opinion    string    `json:"opinion" xorm:"not null comment('意见') VARCHAR(32)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default 'current_timestamp()' TIMESTAMP"`
}

func (*AuditOpinion) TableName() string {
	return "liveearth_primary.audit_opinion"
}
