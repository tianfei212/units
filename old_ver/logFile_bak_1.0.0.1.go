package old_ver

//
//import (
//	"bufio"
//	"fmt"
//	"github.com/tianfei212/units/IO"
//	"github.com/tianfei212/units/IO/FileHelpper"
//	"github.com/tianfei212/units/IO/LogHelpper"
//	DayTimeHelpper "github.com/tianfei212/units/Sys/DateTimeHelpper"
//	SortA "github.com/tianfei212/units/Sys/OtherHelpper/SortA/v.1.0"
//	"os"
//	"path"
//	"sort"
//	"strconv"
//	"strings"
//	"time"
//)
//
////////////////////////////
////IO/LogHelpper/logFile_bak_1.0.0.1.go
////author = "Derek Tian"
////Ver = 0.0.0.1
////make time = 3/28/2022 20:49
//// note =
///////////////////////////
//
//var fh FileHelpper.FileH
//
//var ior IO.FWead
//var logFilestr string
//var ch chan interface{}
//var logf1 *os.File
//var logfbr *bufio.Writer
//
//func init() {
//	fh = FileHelpper.FileH{}
//	fmt.Println("here ")
//	ch = make(chan interface{}, 1000)
//	ior = IO.FWead{}
//	//go writeLogFromChan()
//}
//
//// TODO 判断文件是否存在，如果不纯在就创建（包含目录）
//func (f LogHelpper.LoggerFile) createPath(pp string) bool {
//	fmt.Println("检查创建目录")
//	return fh.CreatePath(pp)
//}
//func (f LogHelpper.LoggerFile) createFile(fp string) *os.File {
//	f1, status := fh.OpenFile(fp)
//	fmt.Println("have err")
//	if status {
//		fmt.Println("f1 open")
//		return f1
//	}
//	return nil
//}
//
//func (f LogHelpper.LoggerFile) getFileList(fp, fn string, sortFiredType string) []string {
//	pinfo := fh.GetPathInfo(fp)
//	fli := []string{}
//	mt := make(map[string]string)
//	//fileMap = make(map[string]string)
//	for k, v := range pinfo {
//		if strings.Contains(v.FileName, fn) {
//			//fli = append(fli, k)
//			mt[k] = fmt.Sprintf("%v", v.FileModTime.Unix())
//		}
//	}
//	mc := SortA.MapSort(mt, false, false, sortFiredType)
//
//	for _, v := range mc {
//		fli = append(fli, fmt.Sprintf("%v", v))
//	}
//	fmt.Println(mc)
//	//sort.Strings(fli)
//	//sort.StringsAreSorted(fli)
//	fmt.Println("获取文件并排序")
//	return fli
//}
//
//// TODO 判断是否存在旧的文件，如果有旧的文件将旧的文件改名
//func (f LogHelpper.LoggerFile) StartCheckFile(IsDel bool) {
//	f.createPath(f.FilePath)
//	if IsDel {
//		fn := f.getFileList(f.FilePath, f.FileName, "int64")
//		fd := f.delFile(fn, f.MaxFIleSaveNum)
//		f.changeLogName(fd, f.MaxFIleSaveNum)
//	}
//	FileNmme := f.FileName
//	if f.ByModel == LogHelpper.ByFileTime {
//		fmt.Println("按照时间方式创建文件")
//		filedate := DayTimeHelpper.GetNowStr(f.TimeFormat)
//		FileNmme = fmt.Sprintf("%s_%s", filedate, FileNmme)
//	}
//	logFilestr = path.Join(f.FilePath, FileNmme)
//	fc := ior.Cread(logFilestr)
//	logf1 = fc.Fi
//	logfbr = fc.Write
//	ch <- "........New Log start.........."
//	//chlog = make(chan string, 10)
//
//}
//
//// TODO 删除达到删除条件的文件
//func (f LogHelpper.LoggerFile) delFile(fl []string, i int) []string {
//	fmt.Printf("启动删除：最大保存文件数量：%d，传入文件数量：%d\n", i, len(fl))
//	var vv []string
//	defer func() {
//		if err := recover(); err != nil {
//			panic(err)
//		}
//
//	}()
//	for k, v := range fl {
//		ftmp := fh.RextStr(v)
//		if k > i-1 && ftmp != "LOG" {
//			b := fh.DelFile(v)
//			fmt.Println(v, " del status:", b)
//			if !b {
//				vv = append(vv, v)
//			}
//			fmt.Println("del file :", v)
//			time.Sleep(10)
//		} else {
//			vv = append(vv, v)
//			fmt.Println("no del file :", v)
//		}
//	}
//	return vv
//}
//func (f LogHelpper.LoggerFile) delFileTime(fl []string, i int) {
//	fmt.Printf("最大保存文件数量：%d，传入文件数量：%d\n", i, len(fl))
//	var vv []string
//
//	for k, v := range fl {
//		fmt.Println(k, "|", v)
//		if k > i-1 {
//			b := fh.DelFile(v)
//			fmt.Println(v, " del status:", b)
//			if !b {
//				vv = append(vv, v)
//			}
//			fmt.Println("del file :", v)
//			time.Sleep(10)
//		} else {
//			vv = append(vv, v)
//		}
//	}
//
//}
//
//// TODO 更改目录下的文件名
//
//func (f LogHelpper.LoggerFile) changeLogName(fl []string, MaxFIle int) []string {
//	var nnew []string
//	defer func() {
//		if err := recover(); err != nil {
//			panic(err)
//		}
//	}()
//	if len(fl) > 0 {
//		fmt.Println("目录下存在文件，启动改名，传入list大小：", len(fl))
//		// 有文件
//		// 获得位长
//		format := "%d"
//		//fname:=""
//		le := len(strconv.Itoa(MaxFIle))
//		switch le {
//		case 1:
//			format = "%02d"
//		case 2:
//			format = "%02d"
//		case 3:
//			format = "%03d"
//		case 4:
//			format = "%04d"
//		}
//		if len(fl) > 1 {
//			//有文件
//			for i := len(fl) - 1; i >= 0; i-- {
//				oldName := fl[i]
//				ex := fmt.Sprint(fh.RextStr(oldName))
//				if ex != "LOG" {
//					tmn, err := strconv.Atoi(ex)
//					if err == nil {
//						fname := strings.ReplaceAll(oldName, fmt.Sprintf(".%s", ex), "")
//						tmn++
//						tn := fmt.Sprintf(format, tmn)
//						newName := fmt.Sprintf("%s.%s", fname, tn)
//						if fh.MoveFile(oldName, newName) {
//							fmt.Printf("file name [%s]-][%s]\n", oldName, newName)
//							nnew = append(nnew, newName)
//						} else {
//							time.Sleep(500 * time.Microsecond)
//							fmt.Println("改名失败，等候500mSec")
//						}
//
//						time.Sleep(10)
//					}
//				} else {
//					tn := fmt.Sprintf(format, 1)
//					newName := fmt.Sprintf("%s.%s", oldName, tn)
//					if fh.MoveFile(oldName, newName) {
//						fmt.Printf("file name [%s]-][%s]\n", oldName, newName)
//						nnew = append(nnew, newName)
//					} else {
//						time.Sleep(500 * time.Microsecond)
//						fmt.Println("改名失败，等候500mSec")
//					}
//					fmt.Printf("file name [%s]-][%s]\n", oldName, newName)
//					nnew = append(nnew, newName)
//
//					time.Sleep(10)
//				}
//
//			}
//
//		} else {
//			fmt.Println("只有一个文件！")
//			tn := fmt.Sprintf(format, 1)
//			newName := fmt.Sprintf("%s.%s", fl[0], tn)
//			if fh.MoveFile(fl[0], newName) {
//				fmt.Printf("file name [%s]-][%s]\n", fl[0], newName)
//				nnew = append(nnew, newName)
//			} else {
//				fmt.Println("改名失败，等候500mSec")
//				time.Sleep(500 * time.Microsecond)
//			}
//			fmt.Printf("one file name [%s]-][%s]\n", fl[0], newName)
//			nnew = append(nnew, newName)
//			time.Sleep(10)
//		}
//
//	}
//	sort.Strings(nnew)
//	fmt.Printf("文件改名并返回新列表,[%d]\n", len(nnew))
//	return nnew
//}
//func (f LogHelpper.LoggerFile) spiltFile(fp string) {
//	finfo := fh.FileInfoRStr(fp)
//	switch f.ByModel {
//	case LogHelpper.BYFileRow:
//		if finfo.FIleRows >= f.MaxFileRows {
//			//关闭现在的文件激活更名和删除的动作，并更新文件指针和名称
//			fmt.Println("启动行数切割！！！！")
//			err := logfbr.Flush()
//
//			if err != nil {
//				return
//			}
//			if clstatus := fh.Close(logf1); clstatus {
//				fn := f.getFileList(f.FilePath, f.FileName, "string")
//				if len(fn) > 0 {
//					fd := f.delFile(fn, f.MaxFIleSaveNum)
//					if len(fd) > 0 {
//						f.changeLogName(fd, f.MaxFIleSaveNum)
//
//					}
//				}
//			}
//			f.StartCheckFile(false)
//
//		}
//	case LogHelpper.BYFileSize:
//		//fmt.Printf("filesize[%d]:[%d]\n", finfo.FileSizeByte, f.MaxFIleSize*1024*1024)
//		if uint64(finfo.FileSizeByte) >= f.MaxFIleSize*1024*1024 {
//			fmt.Println("启动文件大小切割！！！！")
//			err := logfbr.Flush()
//
//			if err != nil {
//				return
//			}
//			if fh.Close(logf1) {
//				fmt.Println("关闭文件正常")
//				fn := f.getFileList(f.FilePath, f.FileName, "string")
//				fd := f.delFile(fn, f.MaxFIleSaveNum)
//				f.changeLogName(fd, f.MaxFIleSaveNum)
//			}
//
//			f.StartCheckFile(false)
//		}
//	case LogHelpper.ByFileTime:
//		filtime := strings.Split(path.Base(fp), "_")
//		setTime := DayTimeHelpper.GetNowStr(f.TimeFormat)
//		if setTime != filtime[0] {
//			fmt.Println("启动时间切割！！！！")
//
//			err := logfbr.Flush()
//
//			if err != nil {
//				return
//			}
//			fh.Close(logf1)
//			fn := f.getFileList(f.FilePath, f.FileName, "int64")
//			f.delFileTime(fn, f.MaxFIleSaveNum)
//			//f.changeLogName(fd, f.MaxFIleSaveNum)
//			f.StartCheckFile(false)
//		}
//	}
//	//fmt.Println(finfo)
//}
//
//func writeLogFromChan() {
//	fmt.Printf("启动写的通道\n")
//	defer func() {
//		err := recover()
//		if err != nil {
//			panic(err)
//		} else {
//			fmt.Println("写结束")
//		}
//	}()
//	//var sa string
//	//select {
//	//case sa = <-ch: //读chan
//	//	fmt.Printf("exec success[%s]\n", sa)
//	//	//return
//	//case <-time.After(1 * time.Second):
//	//	fmt.Printf("exec timeout\n")
//	////return
//	//default:
//	//	fmt.Println("no ddata")
//	//}
//	//fmt.Println(len(chlog))
//	for w1 := range ch {
//		//	fmt.Println("a:", w1)
//
//		_, err := logfbr.WriteString(fmt.Sprintln(w1))
//		if err != nil {
//			panic(err)
//		}
//		err = logfbr.Flush()
//		if err != nil {
//			panic(err)
//		}
//	}
//}
//
//// TODO 实现方法
//
//func (f LogHelpper.LoggerFile) DEBUG(format string, a ...interface{}) {
//	f.spiltFile(logFilestr)
//	s1 := fmt.Sprintf("%s", LogHelpper.mLog(LogHelpper.DEBUG, &f.logger, 3, LogHelpper.conlog(format, a...)))
//
//	ch <- s1
//
//}
//func (f LogHelpper.LoggerFile) TRACE(format string, a ...interface{}) {
//	f.spiltFile(logFilestr)
//	s1 := fmt.Sprintf("%s", LogHelpper.mLog(LogHelpper.TRACE, &f.logger, 3, LogHelpper.conlog(format, a...)))
//
//	ch <- s1
//
//}
//func (f LogHelpper.LoggerFile) INFO(format string, a ...interface{}) {
//	f.spiltFile(logFilestr)
//	s1 := fmt.Sprintf("%s", LogHelpper.mLog(LogHelpper.INFO, &f.logger, 3, LogHelpper.conlog(format, a...)))
//
//	ch <- s1
//
//}
//func (f LogHelpper.LoggerFile) WARNING(format string, a ...interface{}) {
//	f.spiltFile(logFilestr)
//	s1 := fmt.Sprintf("%s", LogHelpper.mLog(LogHelpper.WARNING, &f.logger, 3, LogHelpper.conlog(format, a...)))
//
//	ch <- s1
//
//}
//func (f LogHelpper.LoggerFile) ERROR(format string, a ...interface{}) {
//	f.spiltFile(logFilestr)
//	s1 := fmt.Sprintf("%s", LogHelpper.mLog(LogHelpper.ERROR, &f.logger, 3, LogHelpper.conlog(format, a...)))
//
//	ch <- s1
//
//}
//func (f LogHelpper.LoggerFile) FATAL(format string, a ...interface{}) {
//	f.spiltFile(logFilestr)
//	s1 := fmt.Sprintf("%s", LogHelpper.mLog(LogHelpper.FATAL, &f.logger, 3, LogHelpper.conlog(format, a...)))
//
//	ch <- s1
//
//}
