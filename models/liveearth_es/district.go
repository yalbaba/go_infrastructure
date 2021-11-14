/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/9/2
   Description :
-------------------------------------------------
*/

package liveearth_es

type District struct {
	Id       int    `json:"id"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Level    int    `json:"level"`
}

func (*District) Index() string {
	return "district"
}
