package utils

import (
	"net/url"
	"time"

	"github.com/zlyuancn/zstr"
)

const CDNAuthPrivateKey = "sK2QEU75BxMDoUUw"

var CDN = new(cdnUtil)

type cdnUtil struct{}

func (*cdnUtil) MakeAuthUrl(uri string, expire time.Duration) string {
	const (
		md5TextFormat = `{@uri}-{@timestamp}-0-0-{@private_key}`
		authKeyFormat = `{@timestamp}-0-0-{@md5hash}`
		outFormat     = `{@uri}?auth_key={@auth_key}`
	)

	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	endTime := time.Now().Add(expire).Unix()

	s := zstr.TemplateRender(md5TextFormat, map[string]interface{}{"uri": u.Path, "timestamp": endTime, "private_key": CDNAuthPrivateKey})
	authKey := zstr.TemplateRender(authKeyFormat, map[string]interface{}{"timestamp": endTime, "md5hash": Crypto.Md5(s)})
	return zstr.TemplateRender(outFormat, map[string]interface{}{"uri": uri, "auth_key": authKey})
}
