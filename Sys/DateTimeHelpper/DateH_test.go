package DayTimeHelpper

import (
	"fmt"
	"testing"
	"time"
)

//////////////////////////
//Sys/DateTimeHelpper/DateH_test.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 6/14/2022 10:52
// note =
/////////////////////////

func TestTFormat_Format(t *testing.T) {
	str := "yyyy-mm-dd HH24:mi:ss"
	a := TFormat(str).Format()
	fmt.Println(a)
}

func TestTimeH_StringToTimeStamp(t *testing.T) {
	str := "2006-01-02 15:04:05.000"
	//c, err := time.Parse("2006-01-02 03:04:05", str)
	//fmt.Println(c)
	a, err := TimeH{}.StringToTimeStamp(str, msec)
	fmt.Println(a)
	fmt.Println(err)
}

func TestTimeH_TimeToTimeStamp(t *testing.T) {
	src := time.Now()
	a, err := TimeH{}.TimeToTimeStamp(src, 1)
	fmt.Println(a)
	fmt.Println(err)
	// 需要先初始化timeH
	at := TimeH{TFormat: "yyyy-mm-dd HH24:mi:ss"}
	src1 := at.Now()
	a1, err1 := at.StringToTimeStamp(src1, msec)
	fmt.Println(a1)
	fmt.Println(err1)
}
