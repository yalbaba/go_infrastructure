/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/14
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type AuthorityCompany struct {
	Id            int       `json:"id"`
	CompanyName   string    `json:"company_name"`
	State         int       `json:"state"`
	Deleted       int       `json:"deleted"`
	DeleteTime    time.Time `json:"delete_time"`
	OperatorId    int       `json:"operator_id"`   // 操作人id
	OperatorName  string    `json:"operator_name"` // 操作人名
	DepartmentIds string    `json:"department_ids"`
}

func (*AuthorityCompany) TableName() string {
	return "liveearth_primary.authority_company"
}
