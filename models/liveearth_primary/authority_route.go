/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/13
   Description :
-------------------------------------------------
*/

package liveearth_primary

type AuthorityRoute struct {
	Id         int    `json:"id"`
	RoutePath  string `json:"route_path"` //路由路径
	ResourceId int    `json:"resource_id"`
	FactorId   int    `json:"factor_id"`
	RouteDesc  string `json:"route_desc"`
}

func (*AuthorityRoute) TableName() string {
	return "liveearth_primary.authority_route"
}
