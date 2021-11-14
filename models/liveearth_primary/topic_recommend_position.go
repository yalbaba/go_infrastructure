package liveearth_primary

import (
	"time"
)

type TopicRecommendPosition struct {
	Id                  int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	TopicId             int       `json:"topic_id" xorm:"not null default 0 comment('话题id') INT(10)"`
	Position            int       `json:"position" xorm:"not null comment('位置id') unique TINYINT(4)"`
	RecommendDistrictId int       `json:"recommend_district_id" xorm:"recommend_district_id"`
	AreaName            string    `json:"area_name" xorm:"not null default '' comment('区域名称') VARCHAR(32)"`
	RecommendStart      time.Time `json:"recommend_start" xorm:"comment('推荐开始时间') TIMESTAMP"`
	RecommendEnd        time.Time `json:"recommend_end" xorm:"comment('推荐结束时间') TIMESTAMP"`
	IsOccupy            int       `json:"is_occupy" xorm:"not null default 0 comment('是否占用: 10 = 否; 20 = 是') TINYINT(4)"`
	CreateTime          time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime          time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*TopicRecommendPosition) TableName() string {
	return "liveearth_primary.topic_recommend_position"
}
