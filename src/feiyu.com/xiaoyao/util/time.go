package xyutil

//定义一些常用的操作接口

import (
	"fmt"
	"time"
)

const (
	SecondsPerDay = 60 * 60 * 24
)

// 当前时间 (秒)
func CurTimeSec() int64 {
	return time.Now().Unix()
}

// 当前时间 (纳秒)
func CurTimeNs() int64 {
	return time.Now().UnixNano()
}

// 当前时间 (微秒)
func CurTimeUs() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// 当前时间 (毫秒)
func CurTimeMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 当前时间的字符串格式：yyyymmddhhmmss
func CurTimeStr() (str string) {
	now := time.Now()
	str = fmt.Sprintf("%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	return
}

// 把数值时间转成文本时间
func ToStrTime(timestamp int64) (str string) {
	str = time.Unix(timestamp, 0).Format(time.RFC3339)
	return
}

//计算起止时间戳的间隔天数
// begin 开始时间戳
// end   结束时间戳
//return
// day 天数
func DayDiff(begin, end int64) (day int64) {
	timeBeginTmp := time.Unix(begin, 0)
	timeEndTmp := time.Unix(end, 0)
	timeBegin := time.Date(timeBeginTmp.Year(), timeBeginTmp.Month(), timeBeginTmp.Day(), 0, 0, 0, 0, time.Local).Unix()
	timeEnd := time.Date(timeEndTmp.Year(), timeEndTmp.Month(), timeEndTmp.Day(), 0, 0, 0, 0, time.Local).Unix()
	day = int64((timeEnd - timeBegin) / (24 * 60 * 60))
	return
}

//获取某天的起始时间戳
// dayTime int64 当前时间戳
//return
// begin int64 当天0点的时间戳
// end   int64 当天24点的时间戳
func TimestampRange(dayTime int64) (begin, end int64) {
	year, month, day := time.Unix(dayTime, 0).Date()
	begin = time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	end = begin + int64(SecondsPerDay)
	return
}

//获取某天的起始时间戳
//return
// begin int64 当天0点的时间戳
// end   int64 当天24点的时间戳
func CurTimestampRange() (int64, int64) {
	return TimestampRange(time.Now().Unix())
}
