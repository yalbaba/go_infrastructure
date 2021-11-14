/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/8/20
   Description :
-------------------------------------------------
*/

package liveearth_primary

type FootprintTagRelationship struct {
	Id          int `json:"id"`
	FootprintId int `json:"footprint_id"`
	TagId       int `json:"tag_id"`
	TagLevel    int `json:"tag_level"`
}

func (*FootprintTagRelationship) TableName() string {
	return "liveearth_primary.footprint_tag_relationship"
}
