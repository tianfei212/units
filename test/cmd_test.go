package test

import (
	"fmt"
	cmd "github.com/tianfei212/units/Sys/RunSystemCmd"
	"testing"
)

//////////////////////////
//test/cmd_test.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/18/2022 12:11
// note =
/////////////////////////

func TestCmd(t *testing.T) {
	// this is main function
	b := cmd.OsCommandExec("dir c:")
	fmt.Println(b)
}
