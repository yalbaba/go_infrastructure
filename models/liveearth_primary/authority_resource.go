/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/13
   Description :
-------------------------------------------------
*/

package liveearth_primary

type AuthorityResource struct {
	Id            int    `json:"id"`
	ResourceName  string `json:"resource_name"`
	ResourceDesc  string `json:"resource_desc"`
	FactorIds     string `json:"factor_ids"`     //菜单权限因子
	ParentId      int    `json:"parent_id"`      //菜单父id
	ResourceLevel int    `json:"resource_level"` //菜单等级
}

func (*AuthorityResource) TableName() string {
	return "liveearth_primary.authority_resource"
}
