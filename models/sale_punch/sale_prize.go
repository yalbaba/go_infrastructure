package sale_punch

import (
	"time"
)

type SalePrize struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	LiveMerchantId  int       `json:"live_merchant_id" xorm:"not null comment('商家id') index INT(10)"`
	PrizeName       string    `json:"prize_name" xorm:"not null comment('奖品名称') VARCHAR(16)"`
	Amount          int       `json:"amount" xorm:"not null comment('金额,单位(分)') INT(11)"`
	PrizeType       int       `json:"prize_type" xorm:"not null comment('奖品类型: 10 = 消费券; ....') TINYINT(4)"`
	ExchangeAddress string    `json:"exchange_address" xorm:"not null comment('兑换地址') VARCHAR(32)"`
	ExchangeTime    string    `json:"exchange_time" xorm:"not null comment('兑换时间') VARCHAR(16)"`
	SuitableForUse  string    `json:"suitable_for_use" xorm:"not null comment('适用范围') VARCHAR(16)"`
	Phone           string    `json:"phone" xorm:"not null comment('资讯电话') VARCHAR(12)"`
	LaveNum         int       `json:"lave_num" xorm:"not null comment('剩余数量') INT(11)"`
	ConsumeTotal    int       `json:"consume_total" xorm:"not null comment('消耗总数') INT(11)"`
	CreateTime      time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*SalePrize) TableName() string {
	return "sale_punch.sale_prize"
}
