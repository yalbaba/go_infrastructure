/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/14
   Description :
-------------------------------------------------
*/

package liveearth_primary

import "time"

type AuthorityDepartment struct {
	Id             int       `json:"id"`
	DepartmentName string    `json:"department_name"`
	State          int       `json:"state"`
	Deleted        int       `json:"deleted"`
	DeleteTime     time.Time `json:"delete_time"`
	OperatorId     int       `json:"operator_id"`   //操作人id
	OperatorName   string    `json:"operator_name"` //操作人名
}

func (*AuthorityDepartment) TableName() string {
	return "liveearth_primary.authority_department"
}
