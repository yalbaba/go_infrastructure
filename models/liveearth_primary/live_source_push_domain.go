/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/10/19
   Description :
-------------------------------------------------
*/

package liveearth_primary

type LiveSourcePushDomain struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Domain      string `json:"domain"`
	Remark      string `json:"remark"`
	Deleted     int    `json:"deleted"`
	DeletedTime int    `json:"deleted_time"`
}

func (*LiveSourcePushDomain) TableName() string {
	return "liveearth_primary.live_source_push_domain"
}
