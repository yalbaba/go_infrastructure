/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/22
   Description :
-------------------------------------------------
*/

package liveearth_primary

type AppDownChannel struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	ChannelName string `json:"channel_name" xorm:"not null comment('渠道名') VARCHAR(32)"`
	EnName      string `json:"en_name" xorm:"not null comment('渠道英文名') VARCHAR(32)"`
	State       int    `json:"state" xorm:"not null default 10 comment('开启状态: 10=关闭; 20=开启') TINYINT(3)"`
}

func (*AppDownChannel) TableName() string {
	return "liveearth_primary.app_down_channel"
}
