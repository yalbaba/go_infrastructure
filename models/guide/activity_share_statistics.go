package guide

import "time"

type ActivityShareStatistics struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ClickNum     int64     `json:"click_num" `      // 点击量
	AppPv        int       `json:"app_pv" `         // app端的pv
	AppUv        int       `json:"app_uv" `         // app端的uv
	WebPv        int       `json:"web_pv" `         // h5端的pv
	WebUv        int       `json:"web_uv" `         // h5端的uv
	TakenPortNum int       `json:"taken_port_num" ` // 活跃用户已报名人数
	TakePortNum  int       `json:"take_port_num" `  // 当日已报名人数
	Date         time.Time `json:"date" `           // 统计日期
}

func (m *ActivityShareStatistics) TableName() string {
	return "guide.activity_share_statistics"
}
