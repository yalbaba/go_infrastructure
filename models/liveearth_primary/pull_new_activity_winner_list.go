package liveearth_primary

import "time"

// PullNewActivityWinnerList ...
type PullNewActivityWinnerList struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// PullNewActivityID 拉新活动id
	PullNewActivityId int `json:"pull_new_activity_id"`
	// UserID 用户id
	UserId string `json:"user_id"`
	// PhoneNum 电话号码
	PhoneNum   string    `json:"phone_num"`
	CreateTime time.Time `json:"create_time"`
}

func (*PullNewActivityWinnerList) TableName() string {
	return "liveearth_primary.pull_new_activity_winner_list"
}
