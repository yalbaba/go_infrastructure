/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/11/16
   Description :
-------------------------------------------------
*/

package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTzUtil(t *testing.T) {
	for _, r := range []struct {
		Name  string
		Lon   float64
		Lat   float64
		Diff1 float32
		Diff2 float32
	}{
		{"巴黎", 2.21, 48.41, 2, 1},
		{"瑞士", 8.74, 47.39, 2, 1},
		{"金边", 104.55, 11.33, 7, 7},
		{"开罗", 31.13, 30.3, 2, 2},
		{"维多利亚市", -123.36, 48.42, -7, -8},
		{"卡尔加里", -114.4, 51.2, -6, -7},
		{"温尼伯", -97.8, 49.53, -5, -6},
		{"多伦多", -79.36, 43.69, -4, -5},
		{"哈利法克斯", -63.582648, 44.65, -3, -4},
		{"圣约翰", -66.03, 45.30, -3, -4},
		{"北京", 116.2, 39.56, 8, 8},
	} {
		t.Run(r.Name, func(t *testing.T) {
			loc := TZ.GetTimezoneOfGeo(r.Lon, r.Lat)

			diff1 := TZ.GetTimezoneDiff(loc, time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC))
			diff2 := TZ.GetTimezoneDiff(loc, time.Date(2020, 12, 10, 0, 0, 0, 0, time.UTC))
			require.Equal(t, r.Diff1, diff1, "夏令时时差错误")
			require.Equal(t, r.Diff2, diff2, "时差错误")
		})
	}
}
