package liveearth_primary

import (
	"time"
)

type TaskProgramNotice struct {
	Id              int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId          string    `json:"user_id" xorm:"not null comment('用户id') VARCHAR(32)"`
	LiveSourceEpdId int       `json:"live_source_epd_id" xorm:"not null comment('节目预告id') INT(11)"`
	JobId           int64     `json:"job_id" xorm:"not null default 0 comment('队列任务id') BIGINT(20)"`
	NoticeTime      time.Time `json:"notice_time" xorm:"not null default CURRENT_TIMESTAMP comment('提醒时间') TIMESTAMP"`
	State           int       `json:"state" xorm:"not null comment('状态: 10 = 订阅后等待推送; 20 = 开始推送; 29 = 推送中间状态;30 = 推送成功; 90 = 用户取消订阅') TINYINT(4)"`
	CreateTime      time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime      time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*TaskProgramNotice) TableName() string {
	return "liveearth_primary.task_program_notice"
}
