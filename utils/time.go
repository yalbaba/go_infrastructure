package utils

import (
	"fmt"
	"time"
)

const (
	MinuteStamp   = 60000         // 一分钟的毫秒数
	HourStamp     = 3600000       // 一小时的毫秒数
	DayStamp      = 86400000      // 一天的毫秒数
	Year2100Stamp = 4102416000000 // 2100年的时间戳
)

const (
	DefaultLayout = "2006-01-02 15:04:05"
	DateLayout    = "2006-01-02"
	TimeLayout    = "15:04:05"
	MinuteLayout  = "2006-01-02 15:04"
	HourLayout    = "2006-01-02 15"
)

// 获取当前时间字符串, YYYY-MM-DD hh:mm:ss
func GetTimeText() string {
	return time.Now().Format(DefaultLayout)
}

// 获取当前时间字符串, YYYY-MM-DD
func GetTimeTextYMD() string {
	return time.Now().Format(DateLayout)
}

// 获取当前时间字符串, hh:mm:ss
func GetTimeTextHMS() string {
	return time.Now().Format(TimeLayout)
}

// 获取当前时间字符串, YYYY-MM-DD hh:mm
func GetTimeTextMinute() string {
	return time.Now().Format(MinuteLayout)
}

// 获取当前时间字符串, YYYY-MM-DD hh
func GetTimeTextHour() string {
	return time.Now().Format(HourLayout)
}

// 获取当前毫秒级时间戳
func GetTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// region 转换

// 将时间转为毫秒级时间戳
func TimeToStamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 将时间转为默认格式的字符串
func TimeToText(t time.Time) string {
	return t.Format(DefaultLayout)
}

// 毫秒级时间戳转时间
func StampToTime(stamp int64) time.Time {
	return time.Unix(0, stamp*1e6)
}

// 将毫秒级时间戳转为默认格式的字符串
func StampToText(stamp int64) string {
	return time.Unix(0, stamp*1e6).Format(DefaultLayout)
}

// 将标准时间文本转为时间
func TextToTime(text string) (time.Time, error) {
	return time.ParseInLocation(DefaultLayout, text, time.Local)
}

// 将标准时间文本转为时间
func TextToTimeOfLayout(text, layout string) (time.Time, error) {
	return time.ParseInLocation(layout, text, time.Local)
}

// 将标准时间文本转为毫秒级时间戳
func TextToStamp(text string) (int64, error) {
	return TextToStampOfLayout(text, DefaultLayout)
}

// 将标准时间文本转为毫秒级时间戳
func TextToStampOfLayout(text, layout string) (int64, error) {
	t, e := time.ParseInLocation(layout, text, time.Local)
	if e != nil {
		return 0, e
	}
	return t.UnixNano() / 1e6, nil
}

// endregion

// 获取当天开始时的毫秒级时间戳(0时0分0秒)
func GetNowDayStartStamp() int64 {
	t := time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t.Unix() * 1e3
}

// 获取一个时间戳表示的那一天开始的时间戳
//
// 首先获取传入时间戳的那一天, 再返回那一天开始(0时0分0秒)的时间戳
func GetStartDayStampOfStamp(stamp int64) int64 {
	t := time.Unix(0, stamp*1e6)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t.Unix() * 1e3
}

// 构建时间自定义文本
func MakeTimeCustomText(stamp int64) string {
	now_stamp := GetTimeStamp()
	interval := now_stamp - stamp

	if interval < HourStamp {
		return fmt.Sprintf("%d分钟前", int(interval/MinuteStamp))
	}
	if interval < DayStamp {
		return fmt.Sprintf("%d小时前", int(interval/HourStamp))
	}
	if interval < DayStamp*7 {
		return fmt.Sprintf("%d天前", int(interval/DayStamp))
	}
	if interval < DayStamp*30 {
		return fmt.Sprintf("%d周前", int(interval/DayStamp/7))
	}
	if interval < DayStamp*365 {
		return fmt.Sprintf("%d月前", int(interval/DayStamp/30))
	}
	return fmt.Sprintf("%d年前", int(interval/DayStamp/365))
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(startDate, endDate string) []string {

	if startDate == endDate {
		return []string{startDate}
	}

	var d []string

	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		return d
	}
	if date2.Before(date) {
		return d
	}

	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}

	return d
}

// WeekByDate  获取日期是星期几
func WeekByDate(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	week := int(t.Weekday())
	if week == 0 {
		return 7
	}
	return week
}
