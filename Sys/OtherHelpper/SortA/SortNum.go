package SortA

import (
	"sort"
)

//////////////////////////
//Sys/OtherHelpper/SortA/SortNum.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/2/2022 12:57
// note =
/////////////////////////

type pxInt64 []int64

func (x pxInt64) Len() int {
	return len(x)
}
func (x pxInt64) Less(i, j int) bool {
	return x[i] < x[j]
}
func (x pxInt64) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
func (x pxInt64) tS(d []int64) pxInt64 {
	return d
}

type Int64Sort struct {
}

func (i Int64Sort) Asc(d []int64) {
	tp := pxInt64{}
	b := tp.tS(d)
	sort.Stable(b)
}
func (f Int64Sort) Desc(d []int64) {
	tp := pxInt64{}
	b := tp.tS(d)
	sort.Stable(sort.Reverse(b))
}
