/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/6/30
   Description :
-------------------------------------------------
*/

package liveearth_es

type AppUser struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func (*AppUser) Index() string {
	return "app_user"
}
