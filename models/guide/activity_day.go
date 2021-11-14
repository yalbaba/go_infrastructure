package guide

import "time"

//活动奖品每日领取情况表
type ActivityDay struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	StepId     int       `json:"step_id"`
	ActivityId int       `json:"activity_id" ` // 活动id
	PrizeId    int       `json:"prize_id"`     //奖品id
	PrizeCount int       `json:"prize_count" ` // 奖品总数
	TakenCount int       `json:"taken_count" ` // 当天已领取总数
	Date       time.Time `json:"date" `        // 活动进行日期
}

func (m *ActivityDay) TableName() string {
	return "guide.activity_day"
}
