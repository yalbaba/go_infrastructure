package liveearth_primary

import "time"

// TaskLiveRobotLike ...
type TaskLiveRobotLike struct {
	Id uint64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	// live_room_id
	LiveRoomId uint64 `json:"live_room_id"`
	// 作者id
	UserId string `json:"user_id"`
	// DayOrder 第几天
	DayOrder uint8 `json:"day_order"`
	// LikeNum 当天的点赞数
	LikeNum uint8 `json:"like_num"`
	// FollowNum 当天的关注数
	FollowNum uint8 `json:"follow_num"`
	// TaskDate 任务执行的日期
	TaskDate string `json:"task_date"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
}

func (*TaskLiveRobotLike) TableName() string {
	return "liveearth_primary.task_live_robot_like"
}
