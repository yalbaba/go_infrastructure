package sale_punch

import (
	"time"
)

type SaleVerificationLog struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	SaleMerchantId int       `json:"sale_merchant_id" xorm:"not null comment('商家id') INT(10)"`
	SalePrizeId    int       `json:"sale_prize_id" xorm:"not null comment('奖品id') INT(10)"`
	CheckCode      string    `json:"check_code" xorm:"not null comment('校验码') CHAR(8)"`
	CustomerPhone  string    `json:"customer_phone" xorm:"not null comment('客户手机号') VARCHAR(11)"`
	VerifyUserId   string    `json:"verify_user_id" xorm:"not null comment('核销人员id') VARCHAR(32)"`
	CreateTime     time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*SaleVerificationLog) TableName() string {
	return "sale_punch.sale_verification_log"
}
