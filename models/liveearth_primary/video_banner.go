package liveearth_primary

import "time"

type VideoBanner struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	VideoId   int    `json:"video_id" xorm:"not null comment('视频id') INT(10)"`
	VideoName string `json:"video_name" xorm:"not null comment('视频名称') VARCHAR(16)"`
	VideoDesc string `json:"video_desc" xorm:"not null comment('视频简介') VARCHAR(64)"`
	CoverUrl  string `json:"cover_url" xorm:"not null comment('视频封面') VARCHAR(1024)"`
	// RecommendLevel 推荐等级
	RecommendLevel     uint8 `json:"recommend_level"`
	RecommendLevelTemp uint8 `json:"recommend_level_temp"`
	// RecommendStartTime 推荐开始时间
	RecommendStartTime time.Time `json:"recommend_start_time"`
	// RecommendEndTime 推荐结束时间
	RecommendEndTime time.Time `json:"recommend_end_time"`
	CreateTime       time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	State            int       `json:"state"`
}

func (*VideoBanner) TableName() string {
	return "liveearth_primary.video_banner"
}
