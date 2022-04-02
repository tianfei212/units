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

//func MapSort(m map[string]string, IsAsc, SortKey bool, sortFiredType string) map[string]string {
//	m1 := make(map[string]string, len(m))
//	switch sortFiredType {
//	case "int64":
//		var va []float64
//		//var values []interface{}
//
//		if SortKey {
//			// 对Key排序
//			for k, _ := range m {
//				i, _ := strconv.ParseFloat(k, 64)
//				va = append(va, i)
//			}
//			if IsAsc {
//				sort.Float64s(va)
//				//sort.Slice(va, func(i, j int) bool {
//				//	return false
//				//})
//			} else {
//				//sort.SliceIsSorted(va, func(i, j int) bool {
//				//	return false
//				//})
//				sort.Float64sAreSorted(va)
//			}
//			for _, k := range va {
//				m1[strconv.FormatFloat(k, 'E', -1, 64)] = m[strconv.FormatFloat(k, 'E', -1, 64)]
//			}
//		} else {
//			// 对Value排序
//			mp := make(map[float64]string, len(m))
//			for k, v := range m {
//				i, _ := strconv.ParseFloat(v, 64)
//				va = append(va, i)
//				mp[i] = k
//			}
//			if IsAsc {
//				//sort.Slice(va, func(i, j int) bool {
//				//	return false
//				//})
//				sort.Float64s(va)
//			} else {
//				//sort.SliceIsSorted(va, func(i, j int) bool {
//				//	return false
//				//})
//				sort.Float64sAreSorted(va)
//			}
//			for _, k := range va {
//				st := mp[k]
//				m1[st] = strconv.FormatFloat(k, 'E', -1, 64)
//			}
//		}
//
//	default:
//		var va []string
//		//var values []interface{}
//
//		if SortKey {
//			// 对Key排序
//			for k, _ := range m {
//				va = append(va, k)
//			}
//			if IsAsc {
//				sort.Strings(va)
//				//sort.Slice(va, func(i, j int) bool {
//				//	return false
//				//})
//			} else {
//				//sort.SliceIsSorted(va, func(i, j int) bool {
//				//	return false
//				//})
//				sort.StringsAreSorted(va)
//			}
//			for _, k := range va {
//				m1[k] = m[k]
//			}
//		} else {
//			// 对Value排序
//			mp := make(map[string]string, len(m))
//			for k, v := range m {
//				va = append(va, v)
//				mp[v] = k
//			}
//			if IsAsc {
//				//sort.Slice(va, func(i, j int) bool {
//				//	return false
//				//})
//				sort.Strings(va)
//			} else {
//				//sort.SliceIsSorted(va, func(i, j int) bool {
//				//	return false
//				//})
//				sort.StringsAreSorted(va)
//			}
//			for _, k := range va {
//				st := mp[k]
//				m1[st] = k
//			}
//		}
//	}
//
//	return m1
//}
