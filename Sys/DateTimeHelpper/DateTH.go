package DayTimeHelpper

//package main

import (
	"strings"
	"time"
)

//////////////////////////
//DayTimeHelpper/DateTimeH.go
//author = "Derek Tian"
//Ver = 0.0.0.2
//make time = 3/19/2022 15:59
// note = 常用的日期和时间中的函数和方法
// (0.0.0.2)添加了GetNumToMode函数，作为获取第几周之类的方法
/////////////////////////

// 计算时间偏移量时用到的固定预设值
const (
	Min15  = 15 * time.Minute
	Min30  = 30 * time.Minute
	Min45  = 45 * time.Minute
	Hour   = time.Hour
	Hour12 = 12 * Hour
	Day    = 24 * time.Hour
	Week   = 7 * Day
)

// 统一初始化时支持的预设模式
const (
	StringS int = iota // RunTDStr
	TimeT              //RunTDtime
)

// 计算时间属于范围时用到的情况
const (
	ByDAY int = iota
	ByHOUR
	ByMIN15
	ByMIN30
	ByWeek
)

// TFormat 定义时间类的基础结构体
type TFormat string

// Format 返回golang的时间格式
func (T TFormat) Format() string {
	if T != "" {
		stimeF := strings.ToUpper(string(T))
		stimeF = strings.ReplaceAll(stimeF, "YYYY", "2006")
		stimeF = strings.ReplaceAll(stimeF, "MM", "01")
		stimeF = strings.ReplaceAll(stimeF, "DD", "02")
		if strings.Contains(stimeF, "HH24") {
			stimeF = strings.ReplaceAll(stimeF, "HH24", "03")
		} else {
			stimeF = strings.ReplaceAll(stimeF, "HH", "03")
		}
		stimeF = strings.ReplaceAll(stimeF, "MI", "04")
		stimeF = strings.ReplaceAll(stimeF, ".SSS", "999")
		stimeF = strings.ReplaceAll(stimeF, "SS", "05")
		return stimeF
	} else {
		return "20060102150405"
	}
}

// TimeH 定义基类
type timeH struct {
	tDformat string
}

// 时间函数的基础接口
type timeHer interface {
	Now() string
}
type NewTimeStarer interface {
	timeHer
	NewTime(SrcTime string, IsUTC bool, offset interface{}, offsetType string) time.Time
	CompareDate(sTime time.Time, dTime time.Time) int
	NumOfDateTime(sTime string, Model int) string
}
type NewTimer interface {
	timeHer
	NewTime(SrcTime time.Time, IsUTC bool, offset interface{}, offsetType string) time.Time
	CompareDate(sTime time.Time, dTime time.Time) int
	NumOfDateTime(sTime time.Time, Model int) string
}

// RunTDStr 定义基类初始化的方法
func RunTDStr(Format string) TimeDateStr {

	vf := TFormat(Format)

	return TimeDateStr{timeH{tDformat: vf.Format()}}

}

// RunTDtime 定义基类初始化的方法
func RunTDtime(Format string) TimeDate {

	vf := TFormat(Format)

	return TimeDate{timeH{tDformat: vf.Format()}}

}
func RunTD(Format string, Ti int) interface{} {
	if Ti == 0 {
		return RunTDStr(Format)
	} else {
		return RunTDtime(Format)
	}

}

// Now 基类中的获取当前时间的方法
func (T timeH) Now() string {
	curTime := time.Now()
	return curTime.Format(T.tDformat)
}
