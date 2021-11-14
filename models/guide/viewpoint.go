/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type Viewpoint struct {
	Id             int64     `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	LandscapeId    int       `json:"landscape_id" gorm:"column:landscape_id"`       // 景区id
	ViewpointName  string    `json:"viewpoint_name" gorm:"column:viewpoint_name"`   // 景点名称
	CoverUrl       string    `json:"cover_url" gorm:"column:cover_url"`             // 景点照片
	Addr           string    `json:"addr" gorm:"column:addr"`                       // 景点地址
	ViewpointPoint Point     `json:"viewpoint_point" gorm:"column:viewpoint_point"` // 景点定位
	Altitude       int       `json:"altitude" gorm:"column:altitude"`               // 海拔, 单位米
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *Viewpoint) TableName() string {
	return "guide.viewpoint"
}
