package sale_punch

import (
	"time"
)

type SaleUserPunch struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	MerchantId int       `json:"merchant_id" xorm:"not null comment('商家id') index unique INT(10)"`
	UserId     string    `json:"user_id" xorm:"not null comment('用户id') unique VARCHAR(32)"`
	PunchType  int       `json:"punch_type" xorm:"not null comment('打卡类型 10 = 到店打卡; 20 = 助力打卡') TINYINT(4)"`
	Identifier string    `json:"identifier" xorm:"not null default '' comment('助力打卡时产生的识别码') unique VARCHAR(8)"`
	ImageUrl   string    `json:"image_url" xorm:"not null default '' comment('照片') VARCHAR(1024)"`
	Lon        float64   `json:"lon" xorm:"not null default 0 DOUBLE"`
	Lat        float64   `json:"lat" xorm:"not null default 0 DOUBLE"`
	State      int       `json:"state" xorm:"not null comment('状态: 10 = 到店打卡成功; 20 = 助力打卡中; 30 = 助力打卡成功') TINYINT(4)"`
	Reviews    string    `json:"reviews" xorm:"not null comment('打卡描述') unique VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP unique TIMESTAMP"`
	UpdateTime time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*SaleUserPunch) TableName() string {
	return "sale_punch.sale_user_punch"
}
