package liveearth_primary

import "time"

type TopicVideoRelationship struct {
	Id                   int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	VideoId              int       `json:"video_id" xorm:"not null comment('视频id') index INT(10)"`
	TopicId              int       `json:"topic_id" xorm:"not null comment('话题id') index INT(10)"`
	IsRecommend          int       `json:"is_recommend" xorm:"not null default 0 comment('当前足迹在此话题下是否推荐: 10 = 不推荐; 20 = 推荐 ') TINYINT(4)"`
	IsRecommendTemp      int       `json:"is_recommend_temp" xorm:"not null default 0 comment('用户设置的是否推荐: 10 = 不推荐; 20 = 推荐 ') TINYINT(4)"`
	RecommendStart       time.Time `json:"recommend_start" xorm:"comment('推荐开始时间') TIMESTAMP"`
	RecommendEnd         time.Time `json:"recommend_end" xorm:"comment('推荐结束时间') TIMESTAMP"`
	RecommendIsPerpetual int       `json:"recommend_is_perpetual" xorm:"not null comment('是否永久推荐: 10 = 不是; 20 = 是') TINYINT(4)"`
	HistoryRecommend     int       `json:"history_recommend" xorm:"not null default 0 comment('是否有推荐历史 10 = 否; 20 = 是') TINYINT(4)"`
	AdminRecommend       int       `json:"admin_recommend" xorm:"not null default 0 comment('是否管理员设置的推荐: 10 = 否; 20 = 是') TINYINT(4)"`
	CreateTime           time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime           time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	RecommendType        int       `json:"recommend_type"` //推荐类型（0：普通，10：热门，20：精选，30：置顶）
	Sort                 int       `json:"sort"`
}

func (*TopicVideoRelationship) TableName() string {
	return "liveearth_primary.topic_video_relationship"
}
