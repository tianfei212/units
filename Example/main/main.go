package main

//////////////////////////
//Example/main/main.go
//author = "Derek Tian"
//Ver = 0.0.0.2
//make time = 3/25/2022 15:39
// note =（0.0.0.2）修改了写入配置文件的接收，使用interface来接收不再限定为字符串
/////////////////////////

import (
	"fmt"
	"github.com/tianfei212/units/IO/RwConfig"
)

func main() {
	// this is main function
	c := RwConfig.ConfigF{FilePath: "o:/TEMP/a3.json"}
	//a := c.Read("mysql")
	//for k, v := range a {
	//	fmt.Printf("[%s]=[%v]\n", k, v)
	//}
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m2["cc.1"] = "bb"
	//	m2["cc.2"] = []string{"cc", "dd", "ee"}
	m1["1a.v"] = "ttt"
	m1["za.1.2"] = []string{"cc", "dd", "ee"}
	b := c.Write(m1)
	fmt.Println(b)
}
