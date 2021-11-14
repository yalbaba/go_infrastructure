/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type LandscapeView struct {
	Id               int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"` // 景区id
	LandscapeId      int       `json:"landscape_id" gorm:"column:landscape_id"`                 // 景区id
	LandscapeName    string    `json:"landscape_name" gorm:"column:landscape_name"`             // 景区名称
	LandscapeDesc    string    `json:"landscape_desc" gorm:"column:landscape_desc"`             // 景区描述
	LandscapeQuality string    `json:"landscape_quality" gorm:"column:landscape_quality"`       // 景区资质
	LandscapeFeature string    `json:"landscape_feature" gorm:"column:landscape_feature"`       // 景区特色
	OfficialAddress  string    `json:"official_address"`                                        // 官方地址
	Tips             string    `json:"tips" gorm:"column:tips"`                                 // 小贴士
	CoverUrl         string    `json:"cover_url" gorm:"column:cover_url"`                       // 景区封面图片
	LandscapeType    int8      `json:"landscape_type" gorm:"column:landscape_type"`             // 景区类型
	IntroVoiceUrl    string    `json:"intro_voice_url" gorm:"column:intro_voice_url"`           // 景区介绍语音链接
	ViewpointNum     int       `json:"viewpoint_num" gorm:"column:viewpoint_num"`               // 景点数
	VrUrl            string    `json:"vr_url" gorm:"column:vr_url"`                             // h5链接
	VrVideoId        int64     `json:"vr_video_id" gorm:"column:vr_video_id"`                   // 视频id
	VrVideoName      string    `json:"vr_video_name" gorm:"column:vr_video_name"`               // vr视频名称
	State            int8      `json:"state" gorm:"column:state"`                               // 状态: 10 = 待审核; 20 = 审核通过; 30 = 未通过; 90 = 下架
	PublishState     int8      `json:"publish_state" gorm:"column:publish_state"`               // 发布状态; 10=不发布; 20=发布
	LiveSwitch       int8      `json:"live_switch" gorm:"column:live_switch"`                   // 直播开关: 10 = 关 ; 20 = 开
	CommentSwitch    int8      `json:"comment_switch" gorm:"column:comment_switch"`             // 解说开关: 10 = 关; 20 = 开
	VrSwitch         int8      `json:"vr_switch" gorm:"column:vr_switch"`                       // vr开关：10 = 关；20 = 开
	VrLinkName       string    `json:"vr_link_name" gorm:"column:vr_link_name"`                 // vr H5 链接 名称
	FirstTagId       int       `json:"first_tag_id" gorm:"column:first_tag_id"`                 // 一级tagId
	SecondTagIds     string    `json:"second_tag_ids" gorm:"column:second_tag_ids"`             // 二级tagId列表, 用英文逗号连接
	Deleted          int8      `json:"deleted" gorm:"column:deleted"`                           // 删除标志: 10 = 未删除; 20 = 已经删除
	Operator         string    `json:"operator" gorm:"column:operator"`                         // 操作人
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime       time.Time `json:"update_time" gorm:"column:update_time"`
	FetchTime        time.Time `json:"fetch_time" gorm:"column:fetch_time"`
	IsRecommendVr    int8      `json:"is_recommend_vr" gorm:"column:is_recommend_vr"`   //是否加入vr推荐：10，否，20，是
	IsRecommendHot   int8      `json:"is_recommend_hot" gorm:"column:is_recommend_hot"` //是否加入热门景区推荐：10，否，20，是
	BaseViewNum      int       `json:"base_view_num"`
}

func (m *LandscapeView) TableName() string {
	return "guide.landscape_view"
}
