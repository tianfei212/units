package LogFileSort

//////////////////////////
//Sys/OtherHelpper/SortA/LogFileSort/LogFileSort.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/2/2022 11:42
// note =
////////////////////////

import (
	"github.com/tianfei212/units/IO/FileHelpper"
	"sort"
)

//type fSotrer interface {
//	SortA.SortAer
//}
type pxFileerFileName []FileHelpper.FileMRow
type pxFileerFileModeTime pxFileerFileName

func (x pxFileerFileName) Len() int {
	return len(x)
}
func (x pxFileerFileName) Less(i, j int) bool {
	return x[i].FileName < x[j].FileName
}
func (x pxFileerFileName) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x pxFileerFileModeTime) Len() int {
	return len(x)
}
func (x pxFileerFileModeTime) Less(i, j int) bool {
	return x[i].FileModTime.Before(x[j].FileModTime)
}
func (x pxFileerFileModeTime) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x pxFileerFileModeTime) tS(d []FileHelpper.FileMRow) pxFileerFileModeTime {
	return d
}
func (x pxFileerFileName) tS(d []FileHelpper.FileMRow) pxFileerFileName {
	return d
}

type LogFileFullName struct {
}

func (f LogFileFullName) Asc(d []FileHelpper.FileMRow) {
	tp := pxFileerFileName{}
	b := tp.tS(d)
	sort.Stable(b)
}
func (f LogFileFullName) Desc(d []FileHelpper.FileMRow) {
	tp := pxFileerFileName{}
	b := tp.tS(d)
	sort.Stable(sort.Reverse(b))
}

type LogFileNameModifyTime struct {
}

func (f LogFileNameModifyTime) Asc(d []FileHelpper.FileMRow) {
	tp := pxFileerFileModeTime{}
	b := tp.tS(d)
	sort.Stable(b)
}
func (f LogFileNameModifyTime) Desc(d []FileHelpper.FileMRow) {
	tp := pxFileerFileModeTime{}
	b := tp.tS(d)
	sort.Stable(sort.Reverse(b))
}
