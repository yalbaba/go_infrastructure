/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/13
   Description :
-------------------------------------------------
*/

package liveearth_primary

import "time"

type AuthorityRole struct {
	Id           int       `json:"id" xorm:"autoincr"`
	RoleName     string    `json:"role_name"`
	CreateTime   int       `json:"create_time"`
	UpdateTime   int       `json:"update_time"`
	Deleted      int       `json:"deleted"`
	DeleteTime   time.Time `json:"delete_time"`
	OperatorId   int       `json:"operator_id"`
	OperatorName string    `json:"operator_name"`
	State        int       `json:"state"`
}

func (*AuthorityRole) TableName() string {
	return "liveearth_primary.authority_role"
}
