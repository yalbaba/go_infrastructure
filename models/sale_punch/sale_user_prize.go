package sale_punch

import (
	"time"
)

type SaleUserPrize struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId         string    `json:"user_id" xorm:"not null unique VARCHAR(32)"`
	PhoneNum       string    `json:"phone_num" xorm:"not null comment('用户手机') VARCHAR(11)"`
	SaleMerchantId int       `json:"sale_merchant_id" xorm:"not null comment('商家id') INT(10)"`
	SalePrizeId    int       `json:"sale_prize_id" xorm:"not null comment('奖品id') unique INT(10)"`
	CheckCode      string    `json:"check_code" xorm:"not null comment('校验码') unique CHAR(8)"`
	State          int       `json:"state" xorm:"not null comment('状态: 10 = 未使用; 20 = 已使用') TINYINT(4)"`
	CreateTime     time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP unique TIMESTAMP"`
}

func (*SaleUserPrize) TableName() string {
	return "sale_punch.sale_user_prize"
}
