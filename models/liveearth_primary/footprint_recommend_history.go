/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/29
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type FootprintRecommendHistory struct {
	Id          int `json:"id"`
	FootprintId int `json:"footprint_id"`
	TopicId     int `json:"topic_id"`
	// OperationType 操作类型; 10=取消推荐; 20=推荐
	OperationType int       `json:"operation_type"`
	OperatorName  string    `json:"operator_name"`
	RecommendTime string    `json:"recommend_time"`
	CreateTime    time.Time `json:"create_time"`
}

func (*FootprintRecommendHistory) TableName() string {
	return "liveearth_primary.footprint_recommend_history"
}
