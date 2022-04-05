package test

import (
	"fmt"
	"testing"
	"time"
)

//////////////////////////
//test/log_test.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/5/2022 09:50
// note =
/////////////////////////

import logh "github.com/tianfei212/units/IO/LogHelpper"

func TestLog(t *testing.T) {
	//d1 := D.RunTDStr("")
	logc := logh.NewLogConsole("debug")
	logc.DEBUG("a1")
	nf := logh.FileSet{
		FilePath:       "o:/TMP/logs",
		FileName:       "test.log",
		ByModel:        logh.ByFileTime,
		TimeFormat:     "YYYYMMDDHHMI",
		MaxFIleSize:    1,
		MaxFIleSaveNum: 10,
		MaxFileRows:    1000,
	}
	logf := logh.NewLogFILE("trace", nf)
	logf.StartCheckFile()

	lc := make(chan string, 100)
	logchan := logh.NewLogChan("debug", lc)
	go func() {
		//fmt.Println("start show chan log ")
		for v := range lc {
			fmt.Printf("From chan get %v\n", v)
		}
	}()
	for i := 0; i < 100000; i++ {
		logf.ERROR("这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试这是测试:%d", i)
		logf.INFO("info")
		time.Sleep(1 * time.Second)
		logchan.TRACE("this is from chan:%v", i)
	}

}
