/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import "time"

type LandscapeLive struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)" gorm:"column:id"`
	LandscapeViewId int       `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区id
	ViewpointId     int64     `json:"viewpoint_id" gorm:"column:viewpoint_id"`           // 景点id
	LiveRoomName    string    `json:"live_room_name" gorm:"column:live_room_name"`       // 直播间名称
	LiveRoomId      int64     `json:"live_room_id" gorm:"column:live_room_id"`           // 直播间id
	IsRecommendLive int8      `json:"is_recommend_live"`
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime      time.Time `json:"update_time" gorm:"column:update_time"`
}

func (m *LandscapeLive) TableName() string {
	return "guide.landscape_live"
}
