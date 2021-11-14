/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type VrVideoMatchUrl struct {
	Id          int64     `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	VrVideoName string    `json:"vr_video_name" gorm:"column:vr_video_name"` // vr视频名称
	Url         string    `json:"url" gorm:"column:url"`                     // VR视频 H5链接URL
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`     // 数据创建时间
}

func (m *VrVideoMatchUrl) TableName() string {
	return "guide.vr_video_match_url"
}
