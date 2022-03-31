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
	DayTimeHelpper "github.com/tianfei212/units/Sys/DateTimeHelpper"
	"time"
)

func main() {
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
	//Dat := DayTimeHelpper.DateTrafficTo{}
	//s := Dat.GetNumToMode("m15", "20220328153031")
	//fmt.Println(s)
	//logger test
	cha := make(chan string, 100)

	go func() {
		fmt.Println("start show chan log ")
		for v := range cha {
			fmt.Printf("From chan get %v\n", v)
		}
	}()
	a(cha)
	//Nlog.DEBUG("haha")
}

func a(lc chan string) {
	nf := LogHelpper.FileSet{
		FilePath:       "./logs",
		FileName:       "test.log",
		ByModel:        LogHelpper.ByFileTime,
		TimeFormat:     "YYYYMMDDHHMI",
		MaxFIleSize:    1,
		MaxFIleSaveNum: 3,
		MaxFileRows:    100,
	}
	NlogF := LogHelpper.NewLogFILE("debug", nf)
	NlogF.StartCheckFile(true)
	Nlog := LogHelpper.NewLogConsole("debug")
	Nlog.DEBUG("cc")
	a1 := make(map[string]int)
	a1["cc1"] = 0
	a1["cc2"] = 2
	Nlog.DEBUG("%s %v", "ac", a1)
	Nlog.INFO("info")
	Nlog.ERROR("err")

	NlogC := LogHelpper.NewLogChan("debug", lc)

	for i := 0; i < 10000000000000; i++ {
		Nlog.ERROR("ERR:%d", i)
		time.Sleep(1 * DayTimeHelpper.Min15 / 1400)
		NlogF.ERROR("ERR vsql -h 10.166.35.91 -d vdb -U offline_user -w radcom -F '*' -At -o /OSRM/TMP/ua.csv -f /OSRM/TMP/b.sql:%d", i)
		//NlogF.INFO("info:%d", i)
		//NlogF.ERROR("err:%d", i)
		NlogC.DEBUG("cc%d", i)
		NlogC.INFO("cc%d", i)
	}

	//for i := 0; i <= 100000; i++ {
	//	Nlog.DEBUG("%s:%d", "test", i)
	//	time.Sleep(10 * time.Second)
	//	fmt.Println("wwww")
	//}

}
