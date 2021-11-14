package liveearth_primary

import "time"

// PullNewHelper ...
type PullNewHelper struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// ActivityID 活动id
	ActivityId uint64 `json:"activity_id"`
	// CandidateUserID 候选用户id
	CandidateUserId string `json:"candidate_user_id"`
	// HelperUserID 助力用户id
	HelperUserId string `json:"helper_user_id"`
	// HelperUserAvatar 助力用户头像
	HelperUserAvatar string    `json:"helper_user_avatar"`
	CreateTime       time.Time `json:"create_time"`
}

func (*PullNewHelper) TableName() string {
	return "liveearth_primary.pull_new_helper"
}
