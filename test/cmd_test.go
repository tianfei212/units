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
	//b := cmd.OsCommandExec("dir c:")
	//fmt.Println(b)
	a1 := []string{"-h", "10.166.35.91", "-d", "vdb", "-U", "offline_user", "-w", "radcom", "-c", "select count(*) from omniq.cdr_gi"}
	//bb := fmt.Sprintf("%v", a1)
	//fmt.Printf("%T,%v", bb, bb)
	//b := cmd.OsCommandExec(strings.ReplaceAll(strings.ReplaceAll(bb, "[", ""), "]", ""))
	//fmt.Println(b)

	//cmd := exec.Command("vsql", "-h 10.166.35.91 -d vdb -U offline_user -w radcom -c \"seelct count(*) from omniq.cdr_gi\"")
	//cmd := exec.Command("vsql", a1...)
	//fmt.Println(cmd)
	//var stdout, stderr bytes.Buffer
	//cmd.Stdout = &stdout // 标准输出
	//cmd.Stderr = &stderr // 标准错误
	//err := cmd.Run()
	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("out:n%snerr:n%sn", outStr, errStr)
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %sn", err)
	//}

	//s1, err := cmd.StdoutPipe()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//if err := cmd.Start(); err != nil {
	//	log.Fatal(err)
	//}
	//if opBytes, err := ioutil.ReadAll(s1); err != nil { // 读取输出结果
	//
	//	log.Fatal(err)
	//
	//} else {
	//
	//	log.Println(string(opBytes))
	//
	//}
	//if err := cmd.Wait(); err != nil {
	//	log.Fatal(err)
	//}
	ca := cmd.OsCommandExPar("vsql", a1)
	fmt.Println(ca)
}
