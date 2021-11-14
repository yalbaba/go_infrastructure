/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/3/19
   Description :
-------------------------------------------------
*/

package liveearth_es

type Clocking struct {
	Id     int    `json:"id"`
	UserId string `json:"user_id"`

	ClockingTitle string `json:"clocking_title"`
	MediaType     int    `json:"media_type"`

	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Location struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"location"`
	PositionType int `json:"position_type"`

	FirstTagId       int      `json:"first_tag_id"`
	SecondTagIds     []int    `json:"second_tag_ids"`
	KeywordIds       []int    `json:"keyword_ids"`
	TopicIds         []int    `json:"topic_ids"`
	MatchUserIds     []string `json:"match_user_ids"`
	MatchDistrictIds []int    `json:"match_district_ids"`
	VisibleStartTime string   `json:"visible_start_time"`
	VisibleEndTime   string   `json:"visible_end_time"`

	DataType     int    `json:"data_type"` // 数据类型
	PublishState int    `json:"publish_state"`
	State        int    `json:"state"`
	Deleted      int    `json:"deleted"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	PublishTime  string `json:"publish_time"`
}

func (*Clocking) Index() string {
	return "clocking"
}
