package liveearth_primary

import "time"

type ShareGoods struct {
	Id         int32     `json:"id"`
	SignalName string    `json:"signal_name"`
	UserId     string    `json:"user_id"`
	DeviceId   string    `json:"device_id"`
	GoodsId    int64     `json:"goods_id"`
	ShareTime  time.Time `json:"share_time"`
	UserIp     string    `json:"user_ip"`
}

func (*ShareGoods) TableName() string {
	return "liveearth_primary.share_goods"
}
