package guide

import "time"

type ActivityLuckyUser struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserName    string    `json:"user_name" `    // 用户名
	PhoneNumber string    `json:"phone_number" ` // 用户手机号码
	Gender      int8      `json:"gender" `       // 性别：10男，20女
	Addr        string    `json:"addr" `         // 收货地址
	ActivityId  int       `json:"activity_id" `  // 中奖的活动id
	IsLucky     int8      `json:"is_lucky" `     // 是否中奖：10否，20是
	CreateTime  time.Time `json:"create_time" `  // 创建时间
}

func (m *ActivityLuckyUser) TableName() string {
	return "guide.activity_lucky_user"
}
