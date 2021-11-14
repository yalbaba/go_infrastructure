package sale_punch

import (
	"time"
)

type SaleLotteryProbability struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	SalePrizeId int       `json:"sale_prize_id" xorm:"not null default 0 comment('奖品id,(为 0 时不中奖)') INT(10)"`
	LotteryDate time.Time `json:"lottery_date" xorm:"not null comment('抽奖日期') DATE"`
	StartTime   time.Time `json:"start_time" xorm:"not null default CURRENT_TIMESTAMP comment('开始') TIMESTAMP"`
	EndTime     time.Time `json:"end_time" xorm:"not null default '0000-00-00 00:00:00' comment('结束') TIMESTAMP"`
	Probability int       `json:"probability" xorm:"not null comment('概率') TINYINT(4)"`
}

func (*SaleLotteryProbability) TableName() string {
	return "sale_punch.sale_lottery_probability"
}
