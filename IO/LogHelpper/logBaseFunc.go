package LogHelpper

import (
	"errors"
	"fmt"
	DayTimeHelpper "github.com/tianfei212/units/Sys/DateTimeHelpper"
	"runtime"
	"strings"
)

//////////////////////////
//IO/LogHelpper/logBaseFunc.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/28/2022 20:50
// note =
/////////////////////////

func mLog(lv LogLevel, l *logger, logv int, msg string) interface{} {

	if l.enable(lv) {
		file, funcName, line := getInfo(logv)
		time := DayTimeHelpper.GetNowStr("YYYY-MM-DD HH24:MI:SS.sss")
		//msg := fmt.Sprintf(format, a...)

		if l.IsOutRowNum {
			//fileName := path.Base(file)
			return fmt.Sprintf("%s|%s|%s:%s:%d|-%v", time, getLogString(lv), file, funcName, line, msg)
		} else {
			return fmt.Sprintf("%s|%s|%s:%d|-%v", time, getLogString(lv), funcName, line, msg)
		}
	}
	return nil
}
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOW"
	}

}
func getInfo(n int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller() failed !")
		return "", "", 0
	}
	funcName := runtime.FuncForPC(pc).Name()
	//f1 := fmt.Sprintf("%v:[%v]:%v", file, funcName, line)
	return file, funcName, line
}
func (l logger) enable(level LogLevel) bool {
	//fmt.Printf("[%v][%v]\n", l.Level, level)
	return l.Level <= level
}

// 内部转换日志等级名称到等级
func parselogLevel(s string) (LogLevel, error) {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("UN KNOW Log Level")
		return UNKNOW, err
	}
}
func smalllog(lv LogLevel, l *logger, format string, a ...interface{}) interface{} {
	l1 := *l
	l1.IsOutRowNum = false
	return mLog(lv, &l1, 4, conlog(format, a...))
}

func conlog(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
