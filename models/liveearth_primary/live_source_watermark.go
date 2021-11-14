/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/11/20
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type LiveSourceWatermark struct {
	Id           int       `json:"id"`
	WatermarkUrl string    `json:"watermark_url"`
	Title        string    `json:"title"`
	State        int       `json:"state"` // State 10=关闭;20=开启
	PushDomain   string    `json:"push_domain"`
	CreateTime   time.Time `json:"create_time"`
}

func (*LiveSourceWatermark) TableName() string {
	return "liveearth_primary.live_source_watermark"
}
