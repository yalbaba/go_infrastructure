package guide

import "time"

type ActivityStep struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	StepName        string    `json:"step_name"`
	ActivityId      int       `json:"activity_id" `       // 活动id
	InviteNum       int       `json:"invite_num" `        // 邀请人数
	IsLuckyStandard int8      `json:"is_lucky_standard" ` // 是否是锦鲤报名标准：10否，20是
	PrizeInfo       string    `json:"prize_info"`         //奖品配置:格式：[{"prize_id":1,"prize_count":100},{...}]
	BigPrizeOdds    float32   `json:"big_prize_odds"`     //大红包的中奖概率
	CreateTime      time.Time `json:"create_time" `       // 创建时间
}

func (m *ActivityStep) TableName() string {
	return "guide.activity_step"
}
