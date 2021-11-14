package sale_punch

import (
	"time"
)

type SaleBoostPunch struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	SaleUserPunchId int       `json:"sale_user_punch_id" xorm:"not null comment('用户打卡表id') INT(10)"`
	UserId          string    `json:"user_id" xorm:"not null comment('助力用户') VARCHAR(32)"`
	CreateTime      time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*SaleBoostPunch) TableName() string {
	return "sale_punch.sale_boost_punch"
}
