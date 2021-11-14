package user

import "testing"

func TestGetToken(t *testing.T) {

	t.Log(GetToken(&UserInfo{
		UserId:     "007b4300fc56c5da2f1b9baf3299dcd8",
		IsLogin:    false,
		DeviceData: nil,
	}))

}
