/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/4
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type FootprintLabel struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	LabelName  string    `json:"label_name" xorm:"comment('标签名') VARCHAR(16)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}

func (*FootprintLabel) TableName() string {
	return "liveearth_primary.footprint_label"
}
