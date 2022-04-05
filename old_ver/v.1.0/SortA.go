package v_1_0

import (
	"fmt"
	"sort"
	"strconv"
)

//////////////////////////
//Sys/OtherHelpper/SortA/SortA.go
//author = "Derek Tian"
//Ver =1.0.0.101
//make time = 4/1/2022 20:21
// note =
/////////////////////////

func MapSort(m map[string]string, IsAsc, SortKey bool, sortFiredType string) map[string]string {
	defer func() {
		err := recover()
		if err != nil {
			panic(err)
		}

	}()
	if len(m) == 1 {
		if SortKey {
			return m
		} else {
			mt := make(map[string]string)
			for k, v := range m {
				mt[v] = k
			}
			return mt
		}

	}
	mout := make(map[string]string, len(m))
	switch sortFiredType {
	case "int64":

		va := IntSlice{}
		mp := make(map[int64]string, len(m))
		if SortKey {
			for k, v := range m {
				vtmp, _ := strconv.ParseInt(k, 10, 64)
				va = append(va, vtmp)
				mp[vtmp] = v
			}
		} else {
			for k, v := range m {
				vtmp, _ := strconv.ParseInt(v, 10, 64)
				va = append(va, vtmp)
				mp[vtmp] = k
			}
		}
		if !sort.IsSorted(va) {
			sort.Sort(va)

			for _, k := range va {
				st := mp[k]
				k1 := strconv.FormatInt(k, 10)
				mout[k1] = st
			}
			if !IsAsc {
				moutb := make(map[string]string, len(m))
				for i := len(va) - 1; i >= 0; i-- {
					fmt.Println(i, mp[va[i]])
					//st := mp[i]
					//k1 := strconv.FormatInt(i, 10)
					//moutb[mp[va[i]]] = strconv.FormatInt(va[i], 10)
					moutb[strconv.FormatInt(va[i], 10)] = mp[va[i]]
				}
				return moutb
			}
			return mout
		}
		return m
	case "string":
		var sortList []string
		//var mtemp = make(map[string]string, len(m))
		if SortKey {
			for k := range m {
				sortList = append(sortList, k)
			}

		} else {
			//for k, v := range m {
			//	sortList = append(sortList, v)
			//	mtemp[v] = k
			//}
			for k := range m {
				sortList = append(sortList, k)
			}
		}
		if IsAsc {
			sort.Strings(sortList)

		} else {
			//sort.StringsAreSorted(sortList)
			sort.Strings(sortList)

		}
		for _, vl := range sortList {
			mout[m[vl]] = vl
		}
		return mout
	}
	return nil
}
