package test

import (
	"fmt"
	filehlper "github.com/tianfei212/units/IO/FileHelpper"
	logS "github.com/tianfei212/units/Sys/OtherHelpper/SortA/LogFileSort"
	"strings"
	"testing"
)

//////////////////////////
//test/other_test.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/5/2022 10:54
// note =
/////////////////////////

import (
	"github.com/tianfei212/units/Sys/OtherHelpper"
)

type FileSet struct {
	FilePath       string
	FileName       string
	SpiltFile      bool
	ByModel        int
	MaxFIleSize    uint64 //byte
	MaxFIleSaveNum int
	TimeFormat     string
	MaxFileRows    uint64
}

var fh filehlper.FileH

func getFileInfo(s string) (FS []filehlper.FileMRow) {
	fmt.Println("获取文件列表")
	for k, v := range fh.GetPathInfo(s) {
		//fmt.Println(k)
		//fmt.Println(fh.FileInfoRStr(k))
		if strings.Contains(v.FileName, "test.log") {
			FS = append(FS, fh.FileInfoRStr(k))
		}

	}
	return
}
func TestOther(t *testing.T) {
	fmt.Println(OtherHelpper.IsNum("-2.3"))
	fmt.Println("==========================")
	b := OtherHelpper.GetValueType("cc")
	fmt.Println(b)
	fmt.Println("==========================")
	a := OtherHelpper.NewDurationStr(-1.2, "Hour")
	fmt.Println(a)
	fmt.Println("==========================")
	fmt.Println("排序前。。。。")
	Fs := getFileInfo("o:/tmp/logs")
	fmt.Println(Fs)
	//
	FL := logS.LogFileNameModifyTime{}
	FL.Desc(Fs)
	fmt.Println("排序后。。。。")
	for k, v := range Fs {
		fmt.Printf("[%v]-[%v]\n", k, v)
	}
}
