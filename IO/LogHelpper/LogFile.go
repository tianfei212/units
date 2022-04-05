package LogHelpper

import (
	"fmt"
	"github.com/tianfei212/units/IO"
	filehlper "github.com/tianfei212/units/IO/FileHelpper"
	DayTimeHelpper "github.com/tianfei212/units/Sys/DateTimeHelpper"
	logsort "github.com/tianfei212/units/Sys/OtherHelpper/SortA/LogFileSort"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//////////////////////////
//IO/LogHelpper/LogFile.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/2/2022 13:18
// note =
/////////////////////////

var wriLogChan chan string
var fh filehlper.FileH
var wgwrite sync.WaitGroup
var rwlock sync.RWMutex

//var wo IO.FWead

type wlog struct {
	IO.FWead
	fileName string
}

var wl wlog

func init() {
	//wriLogChan = make(chan string, 100)
	//fmt.Println("chan start")
	fh = filehlper.FileH{}

}

// 从通道接收内容写入日志文件
func wirteLog() {
	fmt.Println("启动写日志！！！！！")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Write %v \n", err)
			panic(err)
		}
	}()
	defer wgwrite.Done()
	defer func() {
		err := wl.Write.Flush()
		if err != nil {
			return
		}
	}()
	fmt.Println("111建立文件写地址:", wl.FWead, "文件名称:", wl.fileName)
	//for {
	//	select {
	//	case num, ok := <-wriLogChan:
	//		if !ok {
	//			fmt.Println("error")
	//			return
	//		}
	//		fmt.Println("======", num)
	//		//每次从channel取值后sleep 1秒，方便我们分析
	//
	//	}
	//}
	//for vlog := range wriLogChan {
	//	fmt.Println(vlog)
	//}
	for vlog := range wriLogChan {
		rwlock.Lock()
		//fmt.Println("From Chan:", vlog)
		_, err := wl.Write.WriteString(fmt.Sprintln(vlog))

		if err != nil {
			fmt.Println("日志写入错误:", err)
		} else {
			err := wl.Write.Flush()
			if err != nil {
				fmt.Println("flast error:", err)
			}
		}
		rwlock.Unlock()
	}
}
func (f LoggerFile) StartCheckFile() {
	fmt.Println("首次启动。。。。。。。")
	b := fh.CreatePath(path.Join(f.FilePath))
	fmt.Println("创建:", b)
	dtime := DayTimeHelpper.RunTDStr(f.TimeFormat)
	b1 := IO.FWead{}
	var Fname string
	if f.ByModel == ByFileTime {
		Fname = fmt.Sprintf("%s_%s", dtime.Now(), f.FileName)
	} else {
		Fname = f.FileName
	}
	wl = wlog{
		FWead:    b1.Cread(path.Join(f.FilePath, Fname)),
		fileName: path.Join(f.FilePath, Fname),
	}
	fmt.Println("建立文件写地址:", wl.FWead, "文件名称:", wl.fileName)
	wriLogChan = make(chan string)
	wgwrite.Add(1)
	go wirteLog()
	//wgwrite.Wait()
}
func (f LoggerFile) End() {
	wriLogChan <- ""
	close(wriLogChan)
}

// Flog 定义日志结构体
type Flog struct {
	LoggerFile
}

func (F Flog) createPath(pp string) bool {
	fmt.Println("检查创建目录")
	return fh.CreatePath(pp)
}

// 获取并返回日志目录下的所有的文件
func (F Flog) getFileInfo() (FS []filehlper.FileMRow) {
	fmt.Println("获取文件列表")
	for k, v := range fh.GetPathInfo(F.FilePath) {
		if strings.Contains(v.FileName, F.FileName) {
			FS = append(FS, fh.FileInfoRStr(k))
		}

	}
	return
}
func (F Flog) CreateF() {
	fmt.Println("日志文件目录创建。。。。。。。")
	dtime := DayTimeHelpper.RunTDStr(F.TimeFormat)
	b := F.createPath(F.FilePath)
	fmt.Println("创建:", b)
	b1 := IO.FWead{}
	var Fname string
	if F.ByModel == ByFileTime {
		Fname = fmt.Sprintf("%s_%s", dtime.Now(), F.FileName)
	} else {
		Fname = F.FileName
	}
	wl = wlog{
		FWead:    b1.Cread(path.Join(F.FilePath, Fname)),
		fileName: path.Join(F.FilePath, Fname),
	}
	fmt.Println("建立文件写地址:", wl.FWead, "文件名称:", wl.fileName)

}

const (
	FILENAME     string = "FILENAME"
	FILEMODIFIED string = "FILEMODIFIED"
)

// 根据传入的文件list信息进行排序并返回文件列表
func (F Flog) sortF(FS *[]filehlper.FileMRow, ByFiled string, IsAsc bool) []string {
	fmt.Println("启动排序")
	var FSL []string
	defer func() {
		err := recover()
		if err != nil {
			panic(err)
		}
	}()
	if len(*FS) == 1 {
		for _, v := range *FS {
			FSL = append(FSL, path.Join(F.FilePath, v.FileName))
		}
		return FSL
	}
	switch ByFiled {
	case FILEMODIFIED:
		c1 := logsort.LogFileNameModifyTime{}
		if IsAsc {
			c1.Asc(*FS)
		} else {
			c1.Desc(*FS)
		}

	case FILENAME:
		c1 := logsort.LogFileFullName{}
		//c1 := logsort.LogFileNameModifyTime{}
		if IsAsc {
			c1.Asc(*FS)
		} else {
			c1.Desc(*FS)
		}
	}

	for _, v := range *FS {
		FSL = append(FSL, path.Join(F.FilePath, v.FileName))
	}
	return FSL
}

//  删除旧文件
func (F Flog) delOldFile(Fl []string) bool {
	dstatus := false
	for k, v := range Fl {
		if k >= F.MaxFIleSaveNum {

			dstatus = fh.DelFile(v)
			fmt.Printf("删除旧文件【%s】 %v\n", v, dstatus)
		} else {
			fmt.Println("删除跳过文件:", v)
		}

	}
	return dstatus
}

// 批量改名字
func (F Flog) rName(FL []string) []string {
	format := "%d"
	le := len(strconv.Itoa(F.MaxFIleSaveNum))
	var b []string
	switch le {
	case 1:
		format = "%02d"
	case 2:
		format = "%02d"
	case 3:
		format = "%03d"
	case 4:
		format = "%04d"
	}
	switch F.ByModel {
	case BYFileSize:
		fallthrough
	case BYFileRow:
		for i := len(FL) - 1; i >= 0; i-- {
			ex := fmt.Sprint(fh.RextStr(FL[i]))
			if ex != "LOG" {
				exN, _ := strconv.Atoi(ex)
				exN++
				ntime := fmt.Sprintf(format, exN)
				tmName := strings.ReplaceAll(FL[i], fmt.Sprintf(".%s", ex), "")
				newName := fmt.Sprintf("%s.%s", tmName, ntime)
				if fh.MoveFile(FL[i], newName) {
					b = append(b, newName)
					fmt.Printf("[%s]:[%s]\n", FL[i], newName)
					time.Sleep(200)
				}

			} else {
				ntime := fmt.Sprintf(format, 1)
				newName := fmt.Sprintf("%s.%s", FL[i], ntime)
				if fh.MoveFile(FL[i], newName) {
					b = append(b, newName)
					fmt.Printf("[%s]:[%s]\n", FL[i], newName)
					time.Sleep(200)
				}
			}
		}

	}
	sort.Strings(b)
	return b
}

func (F Flog) fSpilt(fname string) {
	//fmt.Println("启动文件分隔检查！！！！")
	finfo := fh.FileInfoRStr(fname)
	defer func() {
		err := recover()
		if err != nil {
			panic(err)
		}
	}()
	switch F.ByModel {
	case BYFileSize:

		if uint64(finfo.FileSizeByte/1024/1024) >= F.MaxFIleSize {
			fmt.Println("按照文件大小分割！！！！")
			wriLogChan <- ""
			// 关闭当前文件
			err := wl.Write.Flush()
			if err != nil {
				fmt.Println(nil)
			} else {
				err := wl.Fi.Close()
				if err != nil {
					fmt.Println(err)
				}
			}
			// 激活批量改名和删除旧文件
			Fln := F.getFileInfo()
			FL := F.sortF(&Fln, FILENAME, true)
			FL = F.rName(FL)
			F.delOldFile(FL)
			// 打开一个新的文件
			F.CreateF()
		}
	case BYFileRow:

		if finfo.FIleRows > F.MaxFileRows {
			fmt.Println("按照文件行数分割！！！！")
			wriLogChan <- ""
			// 关闭当前文件
			err := wl.Write.Flush()
			if err != nil {
				fmt.Println(nil)
			} else {
				err := wl.Fi.Close()
				if err != nil {
					fmt.Println(err)
				}
			}
			// 激活批量改名和删除旧文件
			Fln := F.getFileInfo()
			FL := F.sortF(&Fln, FILENAME, true)
			FL = F.rName(FL)
			F.delOldFile(FL)
			// 打开一个新的文件
			F.CreateF()
		}
	case ByFileTime:
		dtime := DayTimeHelpper.RunTDStr(F.TimeFormat)
		natime := dtime.Now()
		fileNameTime := strings.Split(path.Base(wl.fileName), "_")[0]
		if natime != fileNameTime {
			fmt.Println("按照文件时间分割！！！！")
			err := wl.Write.Flush()
			if err != nil {
				fmt.Println(nil)
			} else {
				err := wl.Fi.Close()
				if err != nil {
					fmt.Println(err)
				}
			}
			Fln := F.getFileInfo()
			FL := F.sortF(&Fln, FILEMODIFIED, false)
			//F.rName(FL)
			F.delOldFile(FL)
			// 打开一个新的文件
			F.CreateF()
		}
	}
}
func (F Flog) DEBUG(format string, a ...interface{}) {
	F.fSpilt(wl.fileName)
	if mes := mLog(DEBUG, &F.logger, 3, conlog(format, a...)); mes != nil {
		s1 := fmt.Sprintf("%s", mes)
		//fmt.Println("write input:", s1)
		wriLogChan <- s1
	}

}

func (F Flog) TRACE(format string, a ...interface{}) {
	F.fSpilt(wl.fileName)
	if mes := mLog(TRACE, &F.logger, 3, conlog(format, a...)); mes != nil {
		s1 := fmt.Sprintf("%s", mes)
		//fmt.Println("write input:", s1)
		wriLogChan <- s1
	}
}
func (F Flog) INFO(format string, a ...interface{}) {
	F.fSpilt(wl.fileName)
	if mes := mLog(INFO, &F.logger, 3, conlog(format, a...)); mes != nil {
		s1 := fmt.Sprintf("%s", mes)
		//fmt.Println("write input:", s1)
		wriLogChan <- s1
	}

}
func (F Flog) WARNING(format string, a ...interface{}) {
	F.fSpilt(wl.fileName)
	if mes := mLog(WARNING, &F.logger, 3, conlog(format, a...)); mes != nil {
		s1 := fmt.Sprintf("%s", mes)
		//fmt.Println("write input:", s1)
		wriLogChan <- s1
	}

}
func (F Flog) ERROR(format string, a ...interface{}) {
	F.fSpilt(wl.fileName)
	if mes := mLog(ERROR, &F.logger, 3, conlog(format, a...)); mes != nil {
		s1 := fmt.Sprintf("%s", mes)
		//fmt.Println("write input:", s1)
		wriLogChan <- s1
	}
}
func (F Flog) FATAL(format string, a ...interface{}) {
	F.fSpilt(wl.fileName)
	if mes := mLog(FATAL, &F.logger, 3, conlog(format, a...)); mes != nil {
		s1 := fmt.Sprintf("%s", mes)
		//fmt.Println("write input:", s1)
		wriLogChan <- s1
	}

}
