package liveearth_primary

import "time"

// TaskFootprintRobotLike ...
type TaskFootprintRobotLike struct {
	Id uint64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	// FootprintId 足迹id
	FootprintId uint64 `json:"footprint_id"`
	// 足迹的作者id
	UserId string `json:"user_id"`
	// DayOrder 第几天
	DayOrder uint8 `json:"day_order"`
	// LikeNum 当天的点赞数
	LikeNum uint8 `json:"like_num"`
	// FollowNum 当天的关注数
	FollowNum uint8 `json:"follow_num"`
	// TaskDate 任务执行的日期
	TaskDate string `json:"task_date"`
	// UpdateTime 创建时间
	UpdateTime time.Time `json:"update_time"`
}

func (*TaskFootprintRobotLike) TableName() string {
	return "liveearth_primary.task_footprint_robot_like"
}
