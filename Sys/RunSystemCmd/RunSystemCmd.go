package RunSystemCmd

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
)

//////////////////////////
//Sys/RunSystemCmd/RunSystemCmd.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 4/18/2022 12:10
// note =
/////////////////////////

type charset string

const (
	UTF8    = charset("UTF-8")
	GB18030 = charset("GB18030")
)

func convertByte2String(byte []byte, charset charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
func OsCommandExec(minglingAndCanshu string) string {
	var zifuji charset
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		zifuji = GB18030
		cmd = exec.Command("cmd", "/C", minglingAndCanshu)
	} else {
		zifuji = UTF8
		cmd = exec.Command("sh", "-c", minglingAndCanshu)
	}

	//Linux调用系统命令： ping两个包，每个最多等待1秒。命令分割用;
	//cmd := exec.Command("sh", "-c", "ping 192.168.99.1 -c 2 -W 1; dir")
	//window调用系统命令：ping两个包，每个最多等待1秒。命令分割用&
	//cmd := exec.Command("cmd", "/C", `ping 192.168.99.1 -n 2 -w 1000 & dir`)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	return fmt.Sprintf("%s\n", convertByte2String(stdoutStderr, zifuji))
}

func OsCommandExPar(Cmd string, P []string) string {
	//var zifuji charset
	var cmd *exec.Cmd
	cmd = exec.Command(Cmd, P...)
	fmt.Println(cmd)
	if stOut, err := cmd.StdoutPipe(); err != nil {
		fmt.Println(err)
	} else {
		//defer stOut.Close()
		if err1 := cmd.Start(); err != nil {
			log.Fatal(err1)
		}
		if opBytes, err := ioutil.ReadAll(stOut); err != nil { // 读取输出结果    对象

			log.Fatal(err)

		} else {

			log.Println("stat:", string(opBytes))
			return string(opBytes)
		}
	}
	return ""
}
