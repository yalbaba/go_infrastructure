/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/10/21
   Description :
-------------------------------------------------
*/

package liveearth_primary

type HotSearchWord struct {
	Id          int    `json:"id"`
	WordType    int    `json:"word_type"` // WordType 词类型; 10=关键词; 20=地点
	Word        string `json:"word"`      // Word 词值
	SortValue   int    `json:"sort_value"`
	IsHot       int    `json:"is_hot"` // 是否热; 10=否; 20=是
	State       int    `json:"state"`  // State 状态; 10=隐藏; 20=显示
	Deleted     int    `json:"deleted"`
	DeletedTime int    `json:"deleted_time"`
}

func (*HotSearchWord) TableName() string {
	return "liveearth_primary.hot_search_word"
}
