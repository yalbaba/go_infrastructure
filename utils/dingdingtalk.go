package utils

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/zlyuancn/zdingtalk"
)

var dingTalkConfig = map[string]struct {
	AccessToken string
	Secret      string
	At          []string
	Title       string
	Text        string
}{
	"push_stream": {
		AccessToken: "3b4ec026d9a38f1099d598ae636e44b1d4b53491fa9f4ed21bc3b009b44dee80",
		Secret:      "SECe521316874723fec129df2364ed8eb918995a122dfffd0e391a307ba5196bae6",
		At:          nil,
		Title:       "推流观察员",
		Text:        "## <font color=#ff0000>服务器资源已不足，请及时添加！[二哈]</font>",
	},
}

//发送钉钉通知
func DingTalk(notifyType string) error {
	msg := zdingtalk.NewMarkdownMsg(dingTalkConfig[notifyType].Title, dingTalkConfig[notifyType].Text)
	msg.AtAll()

	bs, _ := msgpack.Marshal(msg)
	m := new(zdingtalk.Msg)
	_ = msgpack.Unmarshal(bs, m)

	err := zdingtalk.NewDingTalk(dingTalkConfig[notifyType].AccessToken).
		SetSecret(dingTalkConfig[notifyType].Secret).
		Send(msg, 2)
	if err != nil {
		return fmt.Errorf("钉钉通知错误,err:%v", err)
	}

	return nil
}
