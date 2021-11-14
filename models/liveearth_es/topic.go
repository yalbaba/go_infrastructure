/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/6/29
   Description :
-------------------------------------------------
*/

package liveearth_es

type Topic struct {
	TopicName string         `json:"topic_name"`
	State     int            `json:"state"`
	Keywords  []TopicKeyword `json:"keywords"`
}

type TopicKeyword struct {
	Id      int    `json:"id"`
	Keyword string `json:"keyword"`
}

func (*Topic) Index() string {
	return "topic"
}
