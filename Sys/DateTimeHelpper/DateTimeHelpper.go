package DayTimeHelpper

//package main

import (
	"fmt"
	OtherHelpper "github.com/tianfei212/units/Sys/OtherHelpper"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//////////////////////////
//DayTimeHelpper/DateTimeH.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/19/2022 15:59
// note = 常用的日期和时间中的函数和方法
/////////////////////////

const (
	Min15  = 15 * time.Minute
	Min30  = 30 * time.Minute
	Min45  = 45 * time.Minute
	Hour   = time.Hour
	Hour12 = 12 * Hour
	Day    = 24 * time.Hour
	Week   = 7 * Day
)

// 对格式的转换
func timeFormat(i ...interface{}) (F string) {
	/*
		支持YYYY-MM-DD HH24:mi:ss 默认为YYYYMMDDHH24miss
	*/
	var iTime string
	if len(i) > 0 {
		iTime = fmt.Sprint(i[0]) // 通过fmt.Sprint 将传入的内容变成字符格式的部分,只获取第一个部分
	} else {
		iTime = fmt.Sprint(i)
	}

	iTime = strings.ToUpper(iTime)
	switch iTime {
	case "YYYY-MM-DD HH24:MI:SS", "YYYY-MM-DD HH:MI:SS":
		F = "2006-01-02 15:04:05"
	case "YYYY-MM-DD HH24:MI", "YYYY-MM-DD HH:MI":
		F = "2006-01-02 15:04"
	case "YYYYMMDDHHMI":
		F = "200601021504"
	case "YYYY-MM-DD HH24", "YYYY-MM-DD HH":
		F = "2006-01-02 15"
	case "YYYYMMDDHH":
		F = "2006010215"
	case "DD/MM/YYYY HH24:mi":
		F = "02/01/2006 15:04"
	case "YYYY":
		F = "2016"
	case "YYYYMM":
		F = "200601"
	case "YYYYMMDD", "YYYY-MM-DD":
		F = "20060102"
	case "YYYY-MM":
		F = "2006-01"
	default:
		F = "20060102150405"
	}
	return F
}

// GetNowStr 获取当前时间的str格式时间
func GetNowStr() string {
	curTime := time.Now()
	return curTime.Format(timeFormat())
}

// GetNewDate 获取新的时间值
/*
sTime : 支持time.time格式和string格式
timeType： 支持（Day，Hour,WEEK,MIN,S
offset : 偏移量，支持+-值，支持int和float
*/
func GetNewDate(sTime interface{}, timeType string, offset interface{}) time.Time {
	/*
		sTime 为初始时间值
	*/
	var t1 time.Time
	switch OtherHelpper.GetValueType(sTime) {
	case "string":
		t1 = StrToTime(fmt.Sprintf("%v", sTime), false)
	default:

		t1 = sTime.(time.Time)
	}

	var v2 float64
	//v := reflect.ValueOf(offset)
	//b := v.Kind()
	//switch b {
	//case reflect.Int:
	//	//v1, _ = strconv.ParseFloat(fmt.Sprint(v), 64)
	//	v2 = float64(v.Int())
	//case reflect.Float64:
	//	v2 = v.Float()
	//default:
	//	vs := reflect.ValueOf(offset)
	//
	//	oneT, _ := time.ParseDuration(fmt.Sprintf("%v", vs))
	//	return t1.Add(oneT)
	//}
	switch OtherHelpper.GetValueType(offset) {
	case "int":
		v2 = float64(offset.(int))
	case "float64":
		v2 = float64(offset.(float64))
	default:
		vs := reflect.ValueOf(offset)

		oneT, _ := time.ParseDuration(fmt.Sprintf("%v", vs))
		return t1.Add(oneT)
	}
	if timeType == "" {
		timeType = "DAY"
		i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", v2))
		return t1.AddDate(0, 0, i)
	} else {
		oneT, _ := time.ParseDuration(OtherHelpper.NewDurationStr(v2, timeType))
		return t1.Add(oneT)
	}

}

// StrToTime 将输入的字符串格式的值转换为time.Time格式输出
/*
srcTime : 为原始输入值（string）
IsUTC：是否使用UTC时间（ Bool）
TimeType：可选输入
*/
func StrToTime(SrcTime string, IsUTC bool, TimeType ...interface{}) time.Time {
	/*
		srcTime : 为原始输入值（string）
		IsUTC：是否使用UTC时间（ Bool）
		TimeType：可选输入，如果不输入的话将按照SrcTime的长度来进行默认的输入

	*/
	var t string
	if len(TimeType) > 0 {
		t = fmt.Sprint(TimeType[0])
	} else {
		t = fmt.Sprint(TimeType)
	}
	t = strings.ToUpper(t)
	if OtherHelpper.IsNum(SrcTime) || len(TimeType) == 0 {
		switch len(SrcTime) {
		case 4:
			t = "YYYY"
		case 6:
			t = "YYYYMM"
		case 8:
			t = "YYYYMMDD"
		case 10:
			t = "YYYYMMDDHH"
		case 12:
			t = "YYYYMMDDHHMI"
		case 14:
			t = "YYYYMMDDHHMISS"
		}
	}

	if IsUTC {

		a, _ := time.Parse(timeFormat(t), SrcTime)
		return a
	} else {
		a, _ := time.ParseInLocation(timeFormat(t), SrcTime, time.Local)
		return a
	}

}

// CompareDate 两个时间的比较
/*
   如果sTime《dTime ：-1 ，sTime=dTime：0， sTime》dTIme：1
*/
func CompareDate(sTime time.Time, dTime time.Time) int {

	if sTime.Before(dTime) {
		return -1
	}
	if sTime.Equal(dTime) {
		return 0
	}
	if sTime.After(dTime) {
		return 1
	}
	return 0
}
func main() {
	// this is main function
	//	fmt.Println(StrToTime("20220302", false))
	//ss := OtherHelpper.NewDurationStr(-1.2, "Day")
	s1 := GetNewDate("20220318000000", "week", -12.2)

	//s1 := GetNewDate(time.Now(), "", -3*Week)
	fmt.Println(s1.Format(timeFormat()))
}
