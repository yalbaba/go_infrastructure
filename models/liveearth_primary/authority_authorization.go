/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/14
   Description :
-------------------------------------------------
*/

package liveearth_primary

type AuthorityAuthorization struct {
	Id         int `json:"id"`
	RoleId     int `json:"role_id"`
	ResourceId int `json:"resource_id"`
	FactorId   int `json:"factor_id"`
}

func (*AuthorityAuthorization) TableName() string {
	return "liveearth_primary.authority_authorization"
}
