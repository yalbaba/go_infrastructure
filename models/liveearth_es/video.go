/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/10/20
   Description :
-------------------------------------------------
*/

package liveearth_es

type Video struct {
	Id        int    `json:"id"`
	VideoName string `json:"video_name"`
	VideoDesc string `json:"video_desc"`
	Nation    string `json:"nation"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Location  struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"location"`
	Deleted          int      `json:"deleted"`
	IsPublish        int      `json:"is_publish"`
	VerifyState      int      `json:"verify_state"`
	VisibleStartTime string   `json:"visible_start_time"`
	VisibleEndTime   string   `json:"visible_end_time"`
	Top              int      `json:"top"`                // top等级, 需要定时刷入
	RecommendLevel   int      `json:"recommend_level"`    // 推荐等级, 需要定时刷入
	TodayPublishTime string   `json:"today_publish_time"` // 当天发布时间
	ViewNum          int      `json:"view_num"`           // 查看数, 需要轮训刷入
	MatchUserIds     []string `json:"match_user_ids"`
	MatchDistrictIds []int    `json:"match_district_ids"`
	DataType         int      `json:"data_type"` // 数据类型
	FirstTagId       []int    `json:"first_tag_id"`
	SecondTagIds     []int    `json:"second_tag_ids"`
	UserId           string   `json:"user_id"`
	PublishTime      string   `json:"publish_time"`
	MixtureSortTime  string   `json:"mixture_sort_time"` // 混合排序时间, 发布时间
	CreateTime       string   `json:"create_time"`
	UpdateTime       string   `json:"update_time"`
	VideoType        int      `json:"video_type"`
}

func (*Video) Index() string {
	return "video"
}
