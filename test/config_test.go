package test

import (
	"fmt"
	"testing"
)

//////////////////////////
//test/config_test.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/5/2022 12:52
// note =
/////////////////////////

import rconfig "github.com/tianfei212/units/IO/RwConfig"

func TestConfig(t *testing.T) {
	// this is main function
	fmt.Println("写配置文件")
	b := rconfig.ConfigF{FilePath: "o:/tmp/1.json"}
	cc := make(map[string]interface{})
	cc["ccc1"] = 1
	cc["ccc1.bbb"] = "casdt"
	b.Write(cc)
	fmt.Println("读配置文件")
	b := rconfig.ConfigF{FilePath: "o:/tmp/1.json"}
	cs := b.Read("ccc1")
	fmt.Println(cs)
}
