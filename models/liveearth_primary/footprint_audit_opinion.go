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

type FootprintAuditOpinion struct {
	Id                int       `json:"id" xorm:"not null pk autoincr comment('足迹id') INT(10)"`
	FootprintId       int       `json:"footprint_id" xorm:"not null pk autoincr comment('足迹id') INT(10)"`
	Remark            string    `json:"remark" xorm:"comment('备注') VARCHAR(32)"`
	AuditOpinionIds   string    `json:"audit_opinion_ids" xorm:"comment('审核意见id列表') VARCHAR(128)"`
	AuditOpinionDescs string    `json:"audit_opinion_descs" xorm:"comment('审核意见描述') VARCHAR(2048)"`
	UpdateTime        time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	LabelIds          string    `json:"label_ids" xorm:"comment('标签id列表') JSON"`
	LabelNames        string    `json:"label_names" xorm:"comment('标签名列表') JSON"`
}

func (*FootprintAuditOpinion) TableName() string {
	return "liveearth_primary.footprint_audit_opinion"
}
