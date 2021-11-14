package liveearth_primary

import (
	"time"
)

// LiveSourceEpg ...
type LiveSourceEpg struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// LiveSourceID 直播源id
	LiveSourceId int `json:"live_source_id"`
	// Title 节目标题
	Title string `json:"title"`
	// ProgramDesc 节目介绍
	ProgramDesc string `json:"program_desc"`

	EpgPeopleNum int `json:"epg_people_num"`
	// StartTime 节目开始时间
	StartTime time.Time `json:"start_time"`
	// EndTime 节目结束时间
	EndTime time.Time `json:"end_time"`
	// IsRecommend 是否推荐 10= 否; 20 = 是
	IsRecommend int `json:"is_recommend"`
	// BatchID 批次id
	BatchId string `json:"batch_id"`
	// Policy 时间策略
	Policy string `json:"policy"`
	// Multiple 评分倍数
	Multiple float32 `json:"multiple"`
	// ShowDuration 节目时长
	ShowDuration int `json:"show_duration"`
	// ValidPeriodStart 有效期开始时间
	ValidPeriodStart time.Time `json:"valid_period_start"`
	// ValidPeriodEnd 有效期结束时间
	ValidPeriodEnd time.Time `json:"valid_period_end"`

	Operator string `json:"operator"`

	UpdateTime time.Time `json:"update_time"`

	CoverUrl string `json:"cover_url"`
}

func (*LiveSourceEpg) TableName() string {
	return "liveearth_primary.live_source_epg"
}
