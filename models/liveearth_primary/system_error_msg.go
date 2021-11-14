package liveearth_primary

import "time"

// SystemErrorMsg ...
type SystemErrorMsg struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(10)"`
	// UserId 用户名
	UserId string `json:"user_id"`
	// ItemId 出现错误的对象id
	ItemId uint64 `json:"item_id"`
	// Message 错误信息
	Message string `json:"message"`
	// ErrorType 错误类型: 10 = 直播间错误; 20 = 其他;
	ErrorType  int       `json:"error_type"`
	CreateTime time.Time `json:"create_time"`
}

func (*SystemErrorMsg) TableName() string {
	return "liveearth_primary.system_error_msg"
}
