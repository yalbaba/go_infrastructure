/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type LandscapeImgs struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	LandscapeViewId int       `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区id
	ImgUrl          string    `json:"img_url" gorm:"column:img_url"`                     // 图片地址
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime      time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *LandscapeImgs) TableName() string {
	return "guide.landscape_imgs"
}
