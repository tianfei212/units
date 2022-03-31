package old_ver

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//////////////////////////
//FileHelpper/FileH.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/21/2022 10:10
// note = 这里实现了对于文件操作的基本
/////////////////////////

const (
	CREATE int = os.O_CREATE
	TRUNC  int = os.O_TRUNC
	ROWR   int = os.O_RDWR
	SYNC   int = os.O_SYNC
	EXCL   int = os.O_EXCL
	WONLY  int = os.O_WRONLY
	RONLY  int = os.O_RDONLY
)

type FileH struct {
}

// 创建文件的方法

type FileM struct {
	IsDir        bool
	FileName     string
	FileSizeByte int64
	FileModTime  time.Time
}

func RFile(strF string) (*os.File, bool) {

	file, err := os.Open(strF)
	if err != nil {
		fmt.Printf("[%s] ReadFIle and Return *File have error:%v", "RFile", err)
		return nil, false
	}
	return file, true
}

// CFile 根据传入的文件路径返回一个文件接口的指针（默认创建文件时使用）
func CFile(strF string) (*os.File, bool) {
	var parm os.FileMode
	parm = 0766
	file, err := os.OpenFile(strF, CREATE|TRUNC|ROWR, parm)
	if err != nil {
		fmt.Printf("[%s]  Return *File have error:%v", "CFile", err)
		return nil, false
	}
	return file, true
}

// CFileFleg   根据传入的文件路径返回一个文件接口的指针（可以自定义对文件的操作方法）
func CFileFleg(strF string, flag int, param os.FileMode) (*os.File, bool) {
	file, err := os.OpenFile(strF, flag, param)
	if err != nil {
		fmt.Printf("[%s]  Return *File have error:%v", "CFileFleg", err)
		return nil, false
	}
	return file, true
}

// FileClose 根据传入的文件接口指针关闭文件
func FileClose(f *os.File) bool {
	err := f.Close()
	if err != nil {
		return false
	}
	return true
}

// StatusFIle 根据传入的文件接口指针返回文件的基本信息
func StatusFIle(f *os.File) (os.FileInfo, bool) {
	fi, err := f.Stat()
	if err != nil {
		fmt.Printf("[%s]  Return os.FileInfo have error:%v", "StatusFIle", err)
		return nil, false
	}
	return fi, true
}

// GetPathInfo 根据输入的目录返回目录下文件的的信息，返回的结构体为FIleM类型
func GetPathInfo(rootPath string) map[string]FileM {
	FileMap := make(map[string]FileM, 100)
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		var F FileM
		if info.IsDir() {
			F = FileM{
				IsDir:        info.IsDir(),
				FileName:     "",
				FileSizeByte: 0,
				FileModTime:  info.ModTime(),
			}

		} else {
			F = FileM{
				IsDir:        info.IsDir(),
				FileName:     info.Name(),
				FileSizeByte: info.Size(),
				FileModTime:  info.ModTime(),
			}

		}
		FileMap[path] = F
		//
		//fmt.Println(path) //打印path信息
		//fmt.Println(info.ModTime())
		//fmt.Println(info.Name())

		return nil
	})
	if err != nil {
		fmt.Printf("[%s]  have error:%v", "GetPathInfo", err)
		return nil
	}

	return FileMap
}

// DelFile 删除单一文件
func DelFile(FileAbsPath string) bool {
	err := os.Remove(FileAbsPath)
	if err != nil {
		return false
	}
	return true
}

// DelPath 删除带有文件的目录
func DelPath(Path string) bool {
	err := os.RemoveAll(Path)
	if err != nil {
		return false
	}
	return true
}

// DelFileFromPath 从目录中删除特定的文件
func DelFileFromPath(Path, file string) bool {
	b := true
	F := GetPathInfo(Path)
	for k, v := range F {
		va := strings.ToUpper(v.FileName)
		vb := strings.ToUpper(file)
		if strings.Contains(va, vb) {
			b = DelFile(k)
			if !b {
				fmt.Printf("Del File :[%s] have error", k)
			}
		}
	}
	return b
}

// MoveFile 文件移动/重命名
func MoveFile(SrcP, DrcP string) bool {
	err := os.Rename(SrcP, DrcP)
	if err != nil {
		return false
	}
	return true
}

// FileCopy 文件复制
func FileCopy(Src, drc string) bool {
	in, err := os.Open(Src)
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {

		}
	}(in)
	out, err := os.Create(drc)
	if err != nil {
		return false
	}
	_, err = io.Copy(out, in)
	if err != nil {
		return false
	}
	return true
}

// FileCheck 判断文件是否存在并返回后缀
func FileCheck(src string) (bool, string) {
	_, err := os.Stat(src)
	exists := !os.IsNotExist(err)
	if exists {
		s := path.Ext(src)
		return exists, strings.ToUpper(s)
	} else {
		return false, ""
	}

}

// CreatePath 目录的创建
func CreatePath(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

// RextStr 获取文件的后缀名
func RextStr(s string) interface{} {
	if stu, Ext := FileCheck(s); stu {
		return Ext[1:]
	}
	return nil
}
