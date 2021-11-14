package sale_punch

import (
	"time"
)

type SaleUser struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId           string    `json:"user_id" xorm:"not null unique VARCHAR(32)"`
	LaveLotteryTimes int       `json:"lave_lottery_times" xorm:"not null default 0 comment('剩余抽奖次数') INT(10)"`
	IsLottery        int       `json:"is_lottery" xorm:"not null comment('是否中过奖 10 = 否; 20 = 是') TINYINT(4)"`
	CreateTime       time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*SaleUser) TableName() string {
	return "sale_punch.sale_user"
}
