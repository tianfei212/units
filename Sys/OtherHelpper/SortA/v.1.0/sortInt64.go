package v_1_0

//////////////////////////
//Sys/OtherHelpper/SortA/sortInt64.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/30/2022 22:06
// note =
/////////////////////////

type IntSlice []int64

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}
