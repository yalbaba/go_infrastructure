/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

type LandscapeVrH5 struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	LandscapeName string `json:"landscape_name" gorm:"column:landscape_name"` // 景区名称
	VrH5Url       string `json:"vr_h5_url" gorm:"column:vr_h5_url"`           // vr h5 链接
}

func (m *LandscapeVrH5) TableName() string {
	return "guide.landscape_vr_h5"
}
