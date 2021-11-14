/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type VoiceConvertQueue struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	SceneryId     int64     `json:"scenery_id" gorm:"column:scenery_id"`         // 景物id
	SceneryName   string    `json:"scenery_name" gorm:"column:scenery_name"`     // 景物名称
	LandscapeId   int       `json:"landscape_id" gorm:"column:landscape_id"`     // 景区id
	ViewpointId   int64     `json:"viewpoint_id" gorm:"column:viewpoint_id"`     // 景点id
	State         int8      `json:"state" gorm:"column:state"`                   // 状态: 10 = 等待转换; 20 = 转换中; 30 = 转换成功; 40 = 录入成功; 90 = 转换失败
	OriginalText  string    `json:"original_text" gorm:"column:original_text"`   // 原始文本
	ConvertText   string    `json:"convert_text" gorm:"column:convert_text"`     // 转换文本
	VoiceUrl      string    `json:"voice_url" gorm:"column:voice_url"`           // 语音文件url
	VoiceDuration string    `json:"voice_duration" gorm:"column:voice_duration"` // 语音时长
	BatchId       string    `json:"batch_id" gorm:"column:batch_id"`             // 批次号
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}

func (m *VoiceConvertQueue) TableName() string {
	return "guide.voice_convert_queue"
}
