package liveearth_primary

import (
	"time"
)

type Topic struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	TopicName         string    `json:"topic_name" xorm:"not null comment('话题名称') unique VARCHAR(64)"`
	TopicDesc         string    `json:"topic_desc"` //话题描述
	ArticleNum        int       `json:"article_num" xorm:"default 0 comment('累计ugc文章数量') INT(11)"`
	ExposureNum       int       `json:"exposure_num" xorm:"not null default 0 comment('所有文章曝光数累计') INT(11)"`
	ExposureUser      int       `json:"exposure_user" xorm:"not null default 0 comment('文章曝光的用户去重') INT(11)"`
	RecommendPosition int       `json:"recommend_position" xorm:"default 0 comment('推荐位置 0 = 不推荐') TINYINT(4)"`
	RecommendStart    time.Time `json:"recommend_start" xorm:"comment('推荐开始时间') TIMESTAMP"`
	RecommendEnd      time.Time `json:"recommend_end" xorm:"comment('推荐结束时间') TIMESTAMP"`
	// RecommendDistrictID 推荐的区域id: 0 = 全球
	RecommendDistrictID int       `json:"recommend_district_id" xorm:"recommend_district_id"`
	AreaName            string    `json:"area_name" xorm:"not null default '' comment('区域名称') VARCHAR(64)"`
	State               int       `json:"state" xorm:"not null comment('状态 10 = 用户新建,等待审核 20 = 启用') TINYINT(4)"`
	HistoryHot          int       `json:"history_hot" xorm:"not null default 10 comment('是否历史推荐: 10 = 否; 20 = 是') TINYINT(4)"`
	CreateTime          time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime          time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	Keywords            string    `json:"keywords"`        //关键词数组格式
	BasePeopleNum       int       `json:"base_people_num"` //基础参与人数
	CoverUrl            string    `json:"cover_url"`       //话题封面图
	IsChoose            int       `json:"is_choose"`       //是否选中：10：没有选中，不进入话题广场。20：已选中，进入话题广场展示
}

func (*Topic) TableName() string {
	return "liveearth_primary.topic"
}
