package guide

import "time"

type ActivityBase struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ActivityName string    `json:"activity_name" `  // 活动名称
	BaseTakenNum int       `json:"base_taken_num" ` // 已拆红包基础金额（单位：分）
	BaseLuckyNum int       `json:"base_lucky_num" ` // 锦鲤报名人数基础值
	State        int8      `json:"state"`           // 活动状态：10进行中，20已结束
	LotteryTime  time.Time `json:"lottery_time" `   // 开奖时间
	CreateTime   time.Time `json:"create_time" `    // 创建时间
	UpdateTime   time.Time `json:"update_time" `    // 更新时间
}

func (m *ActivityBase) TableName() string {
	return "guide.activity_base"
}
