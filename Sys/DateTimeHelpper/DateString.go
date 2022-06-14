package DayTimeHelpper

import (
	"fmt"
	"math"
	"strings"
	"time"
)

//////////////////////////
//Sys/DateTimeHelpper/DateString.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/4/2022 21:27
// note =
/////////////////////////

type TimeDateStr struct {
	timeH
}

func (T TimeDateStr) ToTime(SrcTime string, IsUTC bool) time.Time {
	if IsUTC {

		a, _ := time.Parse(T.TFormat, SrcTime)
		return a
	} else {
		a, _ := time.ParseInLocation(T.TFormat, SrcTime, time.Local)
		return a
	}
}

// NewTime
// 输入参数: SrcTime 输入初始时间，字符串格式
//输入参数: IsUTC 布尔值，true，false
//输入参数: offset 时间偏移量，支持函数固定值（Min15，Min30，Min45，Hour，Hour12，Day，Week，一个标准字符串或一个数字可以是+-int或+-folat）
//输入参数: offsetType 当offset不是固定值时有效，如果offset是一个string则无效
func (T TimeDateStr) NewTime(SrcTime string, IsUTC bool, offset interface{}, offsetType string) time.Time {
	c := T.ToTime(SrcTime, IsUTC)
	return getNewTime(c, offset, strings.ToUpper(offsetType))
}

// N15M 快速方法
func (T TimeDateStr) N15M(SrcTime string) time.Time {
	//c := T.ToTime(SrcTime, false)
	return T.NewTime(SrcTime, false, Min15, "")
}

// Min30 快速方法
func (T TimeDateStr) Min30(SrcTime string) time.Time {
	return T.NewTime(SrcTime, false, Min30, "")
}

// N45M 快速方法
func (T TimeDateStr) N45M(SrcTime string) time.Time {
	return T.NewTime(SrcTime, false, Min45, "")
}

// N1Hour 快速方法
func (T TimeDateStr) N1Hour(SrcTime string) time.Time {
	return T.NewTime(SrcTime, false, Hour, "")
}

// N12Hour 快速方法
func (T TimeDateStr) N12Hour(SrcTime string) time.Time {
	return T.NewTime(SrcTime, false, Hour12, "")
}

// N1Week 快速方法
func (T TimeDateStr) N1Week(SrcTime string) time.Time {
	return T.NewTime(SrcTime, false, Week, "")
}

// CompareDate 两个时间值得比较，如果stime>dtime (1)),stime=dtime(0),stime<dtime(-1)
func (T TimeDateStr) CompareDate(s string, d string) int {
	sTime := T.ToTime(s, false)
	dTime := T.ToTime(d, false)
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

// NumOfDateTime 计算一个时间值是一年中的第几周、日、小时、以及15分钟，30分钟
func (T TimeDateStr) NumOfDateTime(sTime string, Model int) string {
	defer func() {
		err := recover()
		if err != nil {
			return
		}
	}()
	var now time.Time
	if sTime != "" {
		d1 := RunTDStr("")
		now = d1.ToTime(sTime, false)
	} else {
		now = time.Now()
	}
	switch Model {
	case ByDAY:
		return fmt.Sprintf("Day_%v", now.YearDay())
	case ByHOUR:
		return fmt.Sprintf("Hour_%v", (now.YearDay()-1)*24+now.Hour())
	case ByMIN15:
		v := (now.YearDay()-1)*24*60 + now.Minute()
		return fmt.Sprintf("M15_%d", v/15)
	case ByMIN30:
		v := (now.YearDay()-1)*24*60 + now.Minute()
		return fmt.Sprintf("M30_%d", v/30)
	case ByWeek:
		day1 := int(math.Ceil(float64(now.UTC().Day())) / 7.0)
		_, day1 = now.ISOWeek()
		return fmt.Sprintf("WEEk_%v", day1)
	}
	return ""
}
