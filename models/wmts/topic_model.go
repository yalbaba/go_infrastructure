/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2021/3/2
   Description :
-------------------------------------------------
*/

package wmts

// topic中的用户行为消息
type UserBehaviorMsg struct {
	T int64 // 秒级时间戳
	X int
	Y int
	Z int
}

// topic中的爬虫任务消息
type SpiderTaskMsg struct {
	X int
	Y int
	Z int
}
