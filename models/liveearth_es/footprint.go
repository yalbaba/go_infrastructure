/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/6/29
   Description :
-------------------------------------------------
*/

package liveearth_es

type Footprint struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Location struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"location"`
	IsPublish        int      `json:"is_publish"`
	State            int      `json:"state"`
	Top              int      `json:"top"`
	RecommendScore   int      `json:"recommend_score"`
	Score            int      `json:"score"`
	MatchUserIds     []string `json:"match_user_ids"`
	MatchDistrictIds []int    `json:"match_district_ids"`
	DataType         int      `json:"data_type"` // 数据类型
	FirstTagId       []int    `json:"first_tag_id"`
	SecondTagIds     []int    `json:"second_tag_ids"`
	TagIds           []int    `json:"tag_ids"`
	CreateTime       string   `json:"create_time"`
	UpdateTime       string   `json:"update_time"`
	UserId           string   `json:"user_id"`
}

func (*Footprint) Index() string {
	return "footprint"
}
