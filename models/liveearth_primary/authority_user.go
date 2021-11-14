/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/13
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type AuthorityUser struct {
	Id                int       `json:"id"`
	UserName          string    `json:"user_name"`     // 用户名
	UserPassword      string    `json:"user_password"` // 口令
	Nickname          string    `json:"nickname"`
	PasswordWatermark string    `json:"password_watermark"`
	RoleIds           string    `json:"role_ids"`
	LockedUser        int       `json:"locked_user"`
	CreateTime        int       `json:"create_time"`
	UpdateTime        int       `json:"update_time"`
	LastLoginTime     int       `json:"last_login_time"`
	OperatorId        int       `json:"operator_id"`
	OperatorName      string    `json:"operator_name"`
	Deleted           int       `json:"deleted"`
	DeleteTime        time.Time `json:"delete_time"`
	PhoneNum          string    `json:"phone_num"`
	DepartmentId      int       `json:"department_id"` //所属部门id
	CompanyId         int       `json:"company_id"`
	UserDesc          string    `json:"user_desc"`
}

func (*AuthorityUser) TableName() string {
	return "liveearth_primary.authority_user"
}
