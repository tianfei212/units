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
	return err == nil
}

func NewDurationStr(val float64, timeType string) string {
	t := strings.ToUpper(timeType)
	var s string
	switch t {
	case "D", "DAY":

		vs := strconv.FormatFloat(val*24, 'f', 1, 64)
		s = fmt.Sprintf("%s%s", vs, "h")
	case "HOUR", "H":
		vs := strconv.FormatFloat(val, 'f', 1, 64)
		s = fmt.Sprintf("%s%s", vs, "h")
	case "WEEK", "W":
		vs := strconv.FormatFloat(val*24*7, 'f', 1, 64)
		s = fmt.Sprintf("%s%s", vs, "h")
	case "MIN", "MI":
		vs := strconv.FormatFloat(val, 'f', 1, 64)
		s = fmt.Sprintf("%s%s", vs, "m")
	default:
		vs := strconv.FormatFloat(val, 'f', 1, 64)
		s = fmt.Sprintf("%s%s", vs, "s")
	}
	return s
}

func GetValueType(a interface{}) string {
	v := reflect.ValueOf(a)
	b := v.Kind()
	//fmt.Println(b.String())
	return b.String()
}
