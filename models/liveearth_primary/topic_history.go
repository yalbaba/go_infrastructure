/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/28
   Description :
-------------------------------------------------
*/

package liveearth_primary

import (
	"time"
)

type TopicHistory struct {
	Id      int `json:"id"`
	TopicId int `json:"topic_id"`
	// OperatorType 操作类型: 10=取消热门; 20=设置热门; 30=修改话题名; 40=管理员新增话题; 50=用户新增话题
	OperationType int `json:"operation_type"`
	// OperatorName 操作人名字
	OperatorName string `json:"operator_name"`
	TopicName    string `json:"topic_name"`
	// RecommendPosition 推荐位置
	RecommendPosition int    `json:"recommend_position"`
	RecommendTime     string `json:"recommend_time"`
	// RecommendArea 推荐区域名
	RecommendDistrict string    `json:"recommend_district"`
	CreateTime        time.Time `json:"create_time"`
}

func (*TopicHistory) TableName() string {
	return "liveearth_primary.topic_history"
}
