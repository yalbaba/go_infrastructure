/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type SceneryComment struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	CommentPkgId    int       `json:"comment_pkg_id" gorm:"column:comment_pkg_id"`       // 解说包id
	LandscapeViewId int       `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区视图id
	LandscapeId     int       `json:"landscape_id" gorm:"column:landscape_id"`           // 景区id
	ViewpointId     int64     `json:"viewpoint_id" gorm:"column:viewpoint_id"`           // 景点id
	SceneryId       int64     `json:"scenery_id" gorm:"column:scenery_id"`               // 景物id
	ViewpointName   string    `json:"viewpoint_name" gorm:"column:viewpoint_name"`       // 景点名称
	ViewpointAddr   string    `json:"viewpoint_addr" gorm:"column:viewpoint_addr"`       // 景点地址
	SceneryName     string    `json:"scenery_name" gorm:"column:scenery_name"`           // 景物名称
	PicUrls         string    `json:"pic_urls" gorm:"column:pic_urls"`                   // 景物图片数组
	VoiceDuration   string    `json:"voice_duration" gorm:"column:voice_duration"`       // 语音时长
	VoiceText       string    `json:"voice_text" gorm:"column:voice_text"`               // 语音文字
	VoiceUrl        string    `json:"voice_url" gorm:"column:voice_url"`                 // 语音url
	ListenNum       int       `json:"listen_num" gorm:"column:listen_num"`               // 收听次数
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime      time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *SceneryComment) TableName() string {
	return "guide.scenery_comment"
}
