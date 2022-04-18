package DayTimeHelpper

import (
	"fmt"
	"github.com/tianfei212/units/Sys/OtherHelpper"
	"reflect"
	"time"
)

//////////////////////////
//Sys/DateTimeHelpper/DateThunits.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/4/2022 15:26
// note =
/////////////////////////

func getNewTime(t1 time.Time, offset interface{}, offsetType string) time.Time {

	fmt.Println(fmt.Sprintf("%v", offset))

	switch offset {
	case Min15, Min30, Min45, Hour, Hour12, Day, Week:
		oneT, _ := time.ParseDuration(fmt.Sprintf("%v", offset))
		t1 = t1.Add(oneT)
		return t1
	default:
		var v2 float64
		switch OtherHelpper.GetValueType(offset) {
		case "int":
			v2 = float64(offset.(int))
		case "float64":
			v2 = float64(offset.(float64))
		case "string":
			vs := reflect.ValueOf(offset)
			oneT, _ := time.ParseDuration(fmt.Sprintf("%v", vs))
			return t1.Add(oneT)
		}
		oneT, _ := time.ParseDuration(OtherHelpper.NewDurationStr(v2, offsetType))
		return t1.Add(oneT)
	}

}
