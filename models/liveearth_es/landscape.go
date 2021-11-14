/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/3/22
   Description :
-------------------------------------------------
*/

package liveearth_es

type Landscape struct {
	Id          int    `json:"id"`
	LandscapeId int    `json:"landscape_id"`
	UserId      string `json:"user_id"`

	LandscapeName string `json:"landscape_name"`

	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Location struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"location"`
	PositionType int `json:"position_type"`

	FirstTagId   int   `json:"first_tag_id"`
	SecondTagIds []int `json:"second_tag_ids"`

	DataType     int    `json:"data_type"` // 数据类型
	PublishState int    `json:"publish_state"`
	State        int    `json:"state"`
	Deleted      int    `json:"deleted"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

func (*Landscape) Index() string {
	return "landscape"
}
