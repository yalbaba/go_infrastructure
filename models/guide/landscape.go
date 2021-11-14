/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type Landscape struct {
	Id             int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"` // 景区id
	LandscapeName  string    `json:"landscape_name" gorm:"column:landscape_name"`             // 景区名称
	Continent      string    `json:"continent" gorm:"column:continent"`                       // 洲
	Nation         string    `json:"nation" gorm:"column:nation"`                             // 国家
	Province       string    `json:"province" gorm:"column:province"`                         // 省
	City           string    `json:"city" gorm:"column:city"`                                 // 市
	County         string    `json:"county" gorm:"column:county"`                             // 区县
	Town           string    `json:"town" gorm:"column:town"`                                 // 镇
	Address        string    `json:"address" gorm:"column:address"`                           // 景区详细地址
	LandscapePoint Point     `json:"landscape_point" gorm:"column:landscape_point"`           // 景区定位点
	PositionType   int8      `json:"position_type" gorm:"column:position_type"`               // 地址类型（10：区域，20：点位）
	IsChina        int8      `json:"is_china" gorm:"column:is_china"`                         // 10 = 国内; 20 = 国外
	Deleted        int8      `json:"deleted" gorm:"column:deleted"`                           // 删除标志: 10 = 未删除; 20 = 已经删除
	Operator       string    `json:"operator" gorm:"column:operator"`                         // 操作人
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time"`
	FetchTime      time.Time `json:"fetch_time" gorm:"column:fetch_time"`
}

func (m *Landscape) TableName() string {
	return "guide.landscape"
}
