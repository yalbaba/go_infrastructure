package guide

import "time"

type ActivityPrizeTakeLog struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	PrizeId    int       `json:"prize_id" `    // 红包id
	UserId     string    `json:"user_id" `     // 领取红包的用户id
	StepId     int       `json:"step_id"`      //领取红包的活动阶段
	ActivityId int       `json:"activity_id" ` // 活动id
	RecordId   int       `json:"record_id"`    //用户正在进行的活动轮次id（activity_record的id）
	CreateTime time.Time `json:"create_time" ` // 创建时间
}

func (m *ActivityPrizeTakeLog) TableName() string {
	return "guide.activity_prize_take_log"
}
