package util

import (
	"fmt"
	"strconv"
	"time"
)

/*
	星期转换
*/
var ConvertWeek = map[int]string{
	1: "一",
	2: "二",
	3: "三",
	4: "四",
	5: "五",
	6: "六",
	7: "日",
}

var LUNAR_INFO = []int{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2,
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977,
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970,
	0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950,
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557,
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0,
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0,
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6,
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570,
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0,
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5,
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930,
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530,
	0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45,
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0}

/*
	将当前时间转换成string
*/
func DateTransformString(time time.Time) string {
	return time.Local().Format("2006-01-02 15:04:05")
}

/*
	获取当前时间[不含日期]
*/
func TimeTransformTimeString(time time.Time) string {
	return time.Local().Format("15:04:05")
}

/*
	获取当前时间[不含时分秒]
*/
func TimeTransformDateString(time time.Time) string {
	return time.Local().Format("2006-01-02")
}

/*
	获取明天时间
*/
func Tomorrow() time.Time {
	t := time.Now()
	t = t.AddDate(0, 0, 1)
	return t
}

/*
	获取昨天时间
*/
func Yesterday() time.Time {
	t := time.Now()
	t = t.AddDate(0, 0, -1)
	return t
}

/*
   计算日期与当前时间的差
*/
func IntervalSinceNow(ts int64) string {
	t := time.Unix(0, ts*int64(time.Millisecond)) //将毫秒转换成时间
	d := (NowTimeStamp() - ts) / 1000             //获取与当前时间的时间差
	if d < 60 {
		return "刚刚"
	} else if d < 3600 {
		return strconv.FormatInt(d/60, 10) + "分钟前"
	} else if d < 86400 {
		return strconv.FormatInt(d/3600, 10) + "小时前"
	} else if TimeStamp(CurrentStartTime(time.Now().Add(-time.Hour*24))) <= ts {
		return "昨天" + TimeTransformTimeString(t)
	} else if TimeStamp(CurrentStartTime(time.Now().Add(-time.Hour*24*2))) <= ts {
		return "前天" + TimeTransformTimeString(t)
	} else {
		return TimeTransformDateString(t)
	}
}

/*
	将当前时间转换成时间戳
*/
func NowTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

/*
	获取时间戳
*/
func TimeStamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

/*
	获取指定时间的当天开始时间
*/
func CurrentStartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

/*
	计算该月总天数
*/
func CurrentMouthDays(year int, month uint) int {
	if (month > 31) || (month < 0) {
		fmt.Println("ERROR MONTH")
	}
	bit := 1 << (16 - month)
	if ((LUNAR_INFO[year-1900] & 0x0FFFF) & bit) == 0 {
		return 29
	} else {
		return 30
	}
}
