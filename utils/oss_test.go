/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package utils

import (
	"testing"
)

func TestGetAliyunOssAccesskeyUtil(t *testing.T) {

	got, err := GetAliyunOssAccesskeyUtil("liveearth-image", "user")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(got)
}
