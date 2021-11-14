/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/6/30
   Description :
-------------------------------------------------
*/

package liveearth_es

type Encyclopedias struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Location struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"location"`
	FirstPublishTime string   `json:"first_publish_time"`
	PublishTime      string   `json:"publish_time"`
	State            int      `json:"state"`
	VisibleStartTime string   `json:"visible_start_time"`
	VisibleEndTime   string   `json:"visible_end_time"`
	MatchUserIds     []string `json:"match_user_ids"`
	MatchDistrictIds []int    `json:"match_district_ids"`
	DataType         int      `json:"data_type"` // 数据类型
	FirstTagId       []int    `json:"first_tag_id"`
	SecondTagIds     []int    `json:"second_tag_ids"`
	TopLevel         int      `json:"top_level"`
	RecommendLevel   int      `json:"recommend_level"`
	SortValue        int      `json:"sort_value"`
	UpdateTime       string   `json:"update_time"`
	UserId           string   `json:"user_id"`
	DetailLayout     int      `json:"detail_layout"`
}

func (*Encyclopedias) Index() string {
	return "encyclopedias"
}
