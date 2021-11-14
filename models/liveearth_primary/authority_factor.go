/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/13
   Description :
-------------------------------------------------
*/

package liveearth_primary

type AuthorityFactor struct {
	Id         int    `json:"id"`
	FactorName string `json:"factor_name"`
	FactorDesc string `json:"factor_desc"`
}

func (*AuthorityFactor) TableName() string {
	return "liveearth_primary.authority_factor"
}
