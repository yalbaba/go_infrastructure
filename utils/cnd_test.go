package utils

import (
	"testing"
	"time"
)

func TestCdnUtil_MakeAuthUrl(t *testing.T) {
	uri := "http://cdn.video.earthonline.com/A5C09260-BCA3-40DD-A0CB-A2253E700C70.png"
	t.Log(CDN.MakeAuthUrl(uri, time.Hour))
}
