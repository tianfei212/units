package FileHelpper

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

type FileM struct {
	IsDir          bool
	FileName       string
	FileSizeByte   int64
	FileModTime    time.Time
	FileCreateTime int64
}
type FileMRow struct {
	FileM
	FIleRows uint64
}
type FileMode uint32
type FileH struct {
}

// OpenFile Create 创建文件
func (f FileH) OpenFile(fp string) (*os.File, bool) {
	return f.OpenFileFalg(fp, CREATE|TRUNC|ROWR, 0644)
}

// OpenFileFalg CreateFalg 创建文件
func (f FileH) OpenFileFalg(fp string, flag int, parm FileMode) (*os.File, bool) {
	defer func() {
		err := recover()
		if err != nil {
			return
		}
	}()
	file, err := os.OpenFile(fp, flag, os.FileMode(parm))
	if err != nil {
		fmt.Printf("[%s]  Return *File have error:%v", "CFileFleg", err)
		return nil, false
	}
	return file, true
}

// Open RFile 读文件打开文件
func (f FileH) Open(strF string) (*os.File, bool) {

	file, err := os.Open(strF)
	if err != nil {
		fmt.Printf("[%s] ReadFIle and Return *File have error:%v", "RFile", err)
		return nil, false
	}
	return file, true
}

// Close 关闭文件
func (f FileH) Close(file *os.File) bool {
	err := file.Close()
	if err != nil {
		return false
	}
	return true
}

// FileStatusStr 获取文件信息（输入为字符串文件路径）
func (f FileH) FileStatusStr(fp string) os.FileInfo {
	f1, _ := f.Open(fp)
	return f.FileStatus(f1)
}

// FileStatus 获取文件信息，文件指针
func (f FileH) FileStatus(f1 *os.File) os.FileInfo {
	defer func() {
		err := recover()
		if err != nil {
			return
		}
	}()

	fi, err := f1.Stat()
	if err != nil {
		return nil
	}
	return fi
}

// GetPathInfo 根据输入的目录返回目录下文件的的信息，返回的结构体为FIleM类型
func (f FileH) GetPathInfo(rootPath string) map[string]FileM {
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
		return nil
	})
	if err != nil {
		fmt.Printf("[%s]  have error:%v", "GetPathInfo", err)
		return nil
	}
	return FileMap
}

// DelFile 删除单一文件
func (f FileH) DelFile(FileAbsPath string) bool {
	err := os.Remove(FileAbsPath)
	if err != nil {
		return false
	}
	return true
}

// DelPath 删除带有文件的目录
func (f FileH) DelPath(Path string) bool {
	err := os.RemoveAll(Path)
	if err != nil {
		return false
	}
	return true
}

// DelFileFromPath 从目录中删除特定的文件
func (f FileH) DelFileFromPath(Path, file string) bool {
	b := true
	F := f.GetPathInfo(Path)
	for k, v := range F {
		va := strings.ToUpper(v.FileName)
		vb := strings.ToUpper(file)
		if strings.Contains(va, vb) {
			b = f.DelFile(k)
			if !b {
				fmt.Printf("Del File :[%s] have error", k)
			}
		}
	}
	return b
}

// MoveFile 文件移动/重命名
func (f FileH) MoveFile(SrcP, DrcP string) bool {
	err := os.Rename(SrcP, DrcP)
	if err != nil {
		return false
	}
	return true
}

// FileCopy 文件复制
func (f FileH) FileCopy(Src, drc string) bool {
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
func (f FileH) FileCheck(src string) (bool, string) {
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
func (f FileH) CreatePath(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

// RextStr 获取文件的后缀名
func (f FileH) RextStr(s string) interface{} {
	if stu, Ext := f.FileCheck(s); stu {
		return strings.ToUpper(Ext[1:])
	}
	return nil
}

// FileInfoRStr 根据传入的文件路径返回带有文件行数的文件信息
func (f FileH) FileInfoRStr(fp string) FileMRow {
	Fsufx := strings.ToUpper(fmt.Sprintf("%v", f.RextStr(fp)))
	Fs := FileMRow{}
	f1, _ := f.Open(fp)

	if Fsufx == "CSV" || Fsufx == "TXT" || Fsufx == "LOG" {
		Fs.FIleRows, _ = lineCounter(f1)
	}

	fin := f.FileStatus(f1)

	Fs.IsDir = fin.IsDir()
	Fs.FileName = fin.Name()
	Fs.FileModTime = fin.ModTime()
	Fs.FileSizeByte = fin.Size()
	f.Close(f1)
	return Fs
}
func (f FileH) FileInfoR(fp string, file *os.File) FileMRow {
	Fsufx := f.RextStr(fp)
	Fs := FileMRow{}

	if Fsufx == "CSV" || Fsufx == "TXT" || Fsufx == "LOG" {
		Fs.FIleRows, _ = lineCounter(file)
	}
	fin := f.FileStatus(file)
	Fs.IsDir = fin.IsDir()
	Fs.FileName = fin.Name()
	Fs.FileModTime = fin.ModTime()
	Fs.FileSizeByte = fin.Size()
	switch runtime.GOOS {
	case "windows":
		wFileSys := fin.Sys()
		fmt.Println(wFileSys)
	case "linux":
	}

	return Fs
}

//===========================================
func lineCounter(r io.Reader) (uint64, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return uint64(count), nil

		case err != nil:
			return 0, err
		}
	}
}
