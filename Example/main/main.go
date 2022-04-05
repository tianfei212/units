package main

//////////////////////////
//Example/main/main.go
//author = "Derek Tian"
//Ver = 0.0.0.2
//make time = 3/25/2022 15:39
// note =（0.0.0.2）修改了写入配置文件的接收，使用interface来接收不再限定为字符串
/////////////////////////

import (
	"fmt"
	"github.com/tianfei212/units/IO/LogHelpper"
	DayH "github.com/tianfei212/units/Sys/DateTimeHelpper"
	"time"
)

func main() {
	c := DayH.RunTD("", DayH.StringS).(DayH.TimeDateStr)
	b := c.Now()
	fmt.Println(b)
	//fmt.Println(TimeH.GetNewDate("2020-10-10 10:10:10", "Day", 10))
	//c := TimeH.TimeStr("2020/10/10 10:10:10")
	//fmt.Println(c.GetFormat())

	// this is main function
	//c := RwConfig.ConfigF{FilePath: "o:/TEMP/a3.json"}
	//a := c.Read("mysql")
	//for k, v := range a {
	//	fmt.Printf("[%s]=[%v]\n", k, v)
	//}
	//m1 := make(map[string]interface{})
	//m2 := make(map[string]interface{})
	//m2["cc.1"] = "bb"
	////	m2["cc.2"] = []string{"cc", "dd", "ee"}
	//m1["1a.v"] = "ttt"
	//m1["za.1.2"] = []string{"cc", "dd", "ee"}
	//b := c.Write(m1)
	//fmt.Println(b)
	// time try
	//Dat := TimeH.DateTrafficTo{}
	//s := Dat.GetNumToMode("m15", "20220328153031")
	//fmt.Println(s)
	//logger test
	//cha := make(chan string, 100)
	//
	//go func() {
	//	//fmt.Println("start show chan log ")
	//	for v := range cha {
	//		fmt.Printf("From chan get %v\n", v)
	//	}
	//}()
	//a(cha)
	//Nlog.DEBUG("haha")
}

//type pxFileer []FileHelpper.FileMRow
//
//func TA() (f pxFileer) {
//	for i := 0; i < 100; i++ {
//		row := FileHelpper.FileMRow{
//			FileM: FileHelpper.FileM{
//				IsDir:          false,
//				FileName:       fmt.Sprintf("%02d_log.log", i),
//				FileSizeByte:   0,
//				FileModTime:    time.Now(),
//				FileCreateTime: 0,
//			},
//			FIleRows: 0,
//		}
//		time.Sleep(10)
//		f = append(f, row)
//	}
//	row := FileHelpper.FileMRow{
//		FileM: FileHelpper.FileM{
//			IsDir:          false,
//			FileName:       fmt.Sprintf("%02d_log.log", 10329848403),
//			FileSizeByte:   0,
//			FileModTime:    time.Now(),
//			FileCreateTime: 0,
//		},
//		FIleRows: 0,
//	}
//	f = append(f, row)
//	return f
//}
func a(lc chan string) {

	//c1 := logsort.LogFileNameModifyTime{}
	//cc := TA()
	//c1.Asc(cc)
	//fmt.Println(reflect.TypeOf(cc))
	//for _, v := range cc {
	//	fmt.Println(v)
	//}
	nf := LogHelpper.FileSet{
		FilePath:       "o:/TMP/logs",
		FileName:       "test.log",
		ByModel:        LogHelpper.BYFileSize,
		TimeFormat:     "YYYYMMDDHHMI",
		MaxFIleSize:    1,
		MaxFIleSaveNum: 10,
		MaxFileRows:    1000,
	}
	NlogF := LogHelpper.NewLogFILE("debug", nf)
	NlogF.StartCheckFile()
	//NlogF.StartCheckFile(true)
	Nlog := LogHelpper.NewLogConsole("debug")
	//Nlog.DEBUG("cc")
	//
	//a1 := make(map[string]int)
	//a1["cc1"] = 0
	//a1["cc2"] = 2
	//Nlog.DEBUG("%s %v", "ac", a1)
	//Nlog.INFO("info")
	//Nlog.ERROR("err")
	//
	NlogC := LogHelpper.NewLogChan("debug", lc)
	//
	for i := 0; i <= 1000000; i++ {
		Nlog.ERROR("ERR:%d", i)
		time.Sleep(900 * time.Microsecond)
		if i == 20 {
			fmt.Println("a", i)
		}
		//fmt.Println("b:", i)
		NlogF.DEBUG("wwwww:%d", i)
		//NlogF.DEBUG("asdf:%d", i)
		//NlogF.DEBUG("c1:%d", i)

		NlogF.ERROR("这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试:%d", i)
		NlogF.TRACE("这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试:%d", i)
		NlogF.INFO("这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试:%d", i)
		//	//NlogF.INFO("info:%d", i)
		//	//NlogF.ERROR("err:%d", i)
		NlogC.DEBUG("cc%d", i)
		NlogC.INFO("cc%d", i)
	}
	NlogF.End()
	//
	////for i := 0; i <= 100000; i++ {
	////	Nlog.DEBUG("%s:%d", "test", i)
	////	time.Sleep(10 * time.Second)
	////	fmt.Println("wwww")
	////}

}
