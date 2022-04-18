package OtherHelpper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//////////////////////////
//OtherHelpper/OtherH.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/19/2022 17:09
// note = 杂项的类和方法
/////////////////////////

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	return true
}

func NewDurationStr(val float64, timeType string) string {
	t := strings.ToUpper(timeType)
	var s string
	switch t {
	case "D", "DAY":
		s = fmt.Sprintf("%v%s", val*24, "h")
	case "HOUR", "H":
		s = fmt.Sprintf("%v%s", val, "h")
	case "WEEK", "W":
		s = fmt.Sprintf("%v%s", val*24*7, "h")
	case "MIN", "MI":
		s = fmt.Sprintf("%v%s", val, "m")
	default:
		s = fmt.Sprintf("%v%s", val, timeType)
	}
	return s
}

func GetValueType(a interface{}) string {
	v := reflect.ValueOf(a)
	b := v.Kind()
	//fmt.Println(b.String())
	return b.String()
}
