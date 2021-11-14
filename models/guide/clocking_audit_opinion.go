package guide

import "time"

type ClockingAuditOpinion struct {
	Id                int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"`
	ClockingId        int       `json:"clocking_id" gorm:"column:clocking_id"`
	Remark            string    `json:"remark" gorm:"column:remark"`                           // 备注
	AuditOpinionIds   string    `json:"audit_opinion_ids" gorm:"column:audit_opinion_ids"`     // 审核意见id列表
	AuditOpinionDescs string    `json:"audit_opinion_descs" gorm:"column:audit_opinion_descs"` // 审核意见描述
	UpdateTime        time.Time `json:"update_time" gorm:"column:update_time"`
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`
}

func (m *ClockingAuditOpinion) TableName() string {
	return "guide.clocking_audit_opinion"
}
