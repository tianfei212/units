package main

//////////////////////////
//Example/main/main.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/25/2022 15:39
// note =
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
	m1 := make(map[string]string)
	m1["1a.v"] = "ttt"
	m1["za"] = "c"
	b := c.Write(m1)
	fmt.Println(b)
}
