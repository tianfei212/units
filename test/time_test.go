package test

import (
	"fmt"
	"testing"
	"time"
)

//////////////////////////
//test/time_test.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/4/2022 15:04
// note =
/////////////////////////
import (
	D "github.com/tianfei212/units/Sys/DateTimeHelpper"
)

func TestSum(t *testing.T) {
	d1 := D.RunTDStr("")
	//c := D.TFormat("yyyy-mm-dd")
	//fmt.Println(c.Format())
	fmt.Println(d1.Now())
	test := "20201010101010"
	fmt.Printf("[%v]--[%v]\n", test, d1.NewTime(test, false, "10h1m", "hour"))
	fmt.Println("====================")
	d2 := D.RunTDtime("yyyy-mm-dd HH24:mi:ss")
	fmt.Printf("[%v]--[%v]\n", time.Now(), d2.NewTime(time.Now(), false, D.Min45, "hour"))
	fmt.Println("====================")
	d3 := D.RunTD("yyyy-mm-dd HH24:mi:ss", D.TimeT).(D.NewTimer)
	fmt.Printf("[%v]--[%v]\n", time.Now(), d3.NewTime(time.Now(), false, "10h1m", "hour"))
	fmt.Println("====================")
	fmt.Println(d3.NumOfDateTime(time.Now(), D.ByWeek))
	fmt.Println("====================")
	fmt.Println(D.TFormat("yyyy-mm-dd HH24:mi:ss").Format())
	fmt.Println(d1.NumOfDateTime("20201010101010", D.ByDAY))

}
