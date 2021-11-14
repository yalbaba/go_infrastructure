/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/6/28
   Description :
-------------------------------------------------
*/

package liveearth_es

type LiveRoom struct {
	Id       int    `json:"id"`
	RoomName string `json:"room_name"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Location struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"location"`
	State            int      `json:"state"`
	Deleted          int      `json:"deleted"`
	MatchUserIds     []string `json:"match_user_ids"`
	MatchDistrictIds []int    `json:"match_district_ids"`
	DataType         int      `json:"data_type"` // 数据类型
	FirstTagId       []int    `json:"first_tag_id"`
	SecondTagIds     []int    `json:"second_tag_ids"`
	UserId           string   `json:"user_id"`
	MixtureSortTime  string   `json:"mixture_sort_time"` // 混合排序时间, 节目单开播时间
	WonderfulVideoId int      `json:"wonderful_video_id"`
	CollectionId     int      `json:"collection_id"`
	SortId           int8     `json:"sort_id"`
}

func (*LiveRoom) Index() string {
	return "live_room"
}
