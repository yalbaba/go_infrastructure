package utils

import (
	"time"

	"gopkg.in/ugjka/go-tz.v2/tz"
)

var TZ = new(tzUtil)

type tzUtil struct{}

// 根据经纬度获取时区
func (*tzUtil) GetTimezoneOfGeo(lon, lat float64, def ...*time.Location) *time.Location {
	zone, _ := tz.GetZone(tz.Point{Lon: lon, Lat: lat})
	if len(zone) == 0 {
		if len(def) > 0 {
			return def[0]
		}
		return time.UTC
	}

	loc, _ := time.LoadLocation(zone[0])
	if loc == nil {
		if len(def) > 0 {
			return def[0]
		}
		return time.UTC
	}
	return loc
}

// 获取时区的时差
func (*tzUtil) GetTimezoneDiff(loc *time.Location, t ...time.Time) float32 {
	var now time.Time
	if len(t) > 0 {
		now = t[0]
		now.In(time.UTC) // 不能写在一行, 否则可能导致入参的数据被改变
	} else {
		now = time.Now().In(time.UTC)
	}
	tLoc := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, loc)
	return float32(now.Unix()-tLoc.Unix()) / 3600 // 时间戳和时间文本是相反的
}

// 获取时区 a 相对于时区 b 的时差
func (*tzUtil) GetTimezoneDiffOfZone(a, b *time.Location, t ...time.Time) float32 {
	var now time.Time
	if len(t) > 0 {
		now = t[0]
		now.In(time.UTC) // 不能写在一行, 否则可能导致入参的数据被改变
	} else {
		now = time.Now().In(time.UTC)
	}
	ta := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, a)
	tb := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, b)
	return float32(tb.Unix()-ta.Unix()) / 3600 // 时间戳和时间文本是相反的
}
