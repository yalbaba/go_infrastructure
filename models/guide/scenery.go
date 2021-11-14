/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type Scenery struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	LandscapeId   int       `json:"landscape_id" gorm:"column:landscape_id"`     // 景区id
	ViewpointId   int64     `json:"viewpoint_id" gorm:"column:viewpoint_id"`     // 景点id
	ViewpointName string    `json:"viewpoint_name" gorm:"column:viewpoint_name"` // 景点名称
	SceneryName   string    `json:"scenery_name" gorm:"column:scenery_name"`     // 景物名称
	Intro         string    `json:"intro" gorm:"column:intro"`                   // 景物介绍
	SceneryPoint  Point     `json:"scenery_point" gorm:"column:scenery_point"`   // 点位
	Addr          string    `json:"addr" gorm:"column:addr"`                     // 地址
	IconUrl       string    `json:"icon_url" gorm:"column:icon_url"`
	PicUrl        string    `json:"pic_url" gorm:"column:pic_url"`   // 首图
	PicUrls       string    `json:"pic_urls" gorm:"column:pic_urls"` // 图片列表
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime    time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *Scenery) TableName() string {
	return "guide.scenery"
}
