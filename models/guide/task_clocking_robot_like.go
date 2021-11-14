/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/3/31
   Description :
-------------------------------------------------
*/

package guide

import (
	"time"
)

type TaskClockingRobotLike struct {
	Id         int       `json:"id"`
	ClockingId int       `json:"clocking_id"` // 打卡id
	UserId     string    `json:"user_id"`     // 视频的用户id
	DayOrder   int8      `json:"day_order"`   // 第几天
	LikeNum    int8      `json:"like_num"`    // 当天的点赞数
	FollowNum  int8      `json:"follow_num"`  // 当天的关注数
	TaskDate   string    `json:"task_date"`   // 任务执行的日期
	CreateTime time.Time `json:"create_time"` // 创建时间
}

func (m *TaskClockingRobotLike) TableName() string {
	return "guide.task_clocking_robot_like"
}
