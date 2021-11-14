package sale_punch

import (
	"time"
)

type SaleMerchant struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	MerchantName    string    `json:"merchant_name" xorm:"not null comment('商家名称') VARCHAR(16)"`
	Icon            string    `json:"icon" xorm:"not null comment('icon') VARCHAR(1024)"`
	MerchantAddress string    `json:"merchant_address" xorm:"not null comment('地址') VARCHAR(64)"`
	Area            string    `json:"area" xorm:"not null comment('区域') VARCHAR(8)"`
	Lon             float64   `json:"lon" xorm:"not null DOUBLE"`
	Lat             float64   `json:"lat" xorm:"not null DOUBLE"`
	Category        int       `json:"category" xorm:"not null comment('商家类别: 10 = 购物潮地 20 = 视听剧苑 30 = 文鉴艺廊 40 = 亲子乐园 50 = 乐动场馆 60 = 晚味去处 70 = 旅游景区 80 = 风情街区 90 = 医美空间 100 = 学习时点') index TINYINT(4)"`
	IsShow          int       `json:"is_show" xorm:"not null default 20 comment('是否展示 10 = 否; 20 = 是') TINYINT(4)"`
	CreateTime      time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdateTime      time.Time `json:"update_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (*SaleMerchant) TableName() string {
	return "sale_punch.sale_merchant"
}
