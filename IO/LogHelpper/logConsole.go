package LogHelpper

import (
	"fmt"
)

//////////////////////////
//IO/LogHelpper/logConsole.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/28/2022 18:58
// note =
/////////////////////////

func (l *LoggerConsole) DEBUG(format string, a ...interface{}) {
	if msg := mLog(DEBUG, &l.logger, 3, conlog(format, a...)); msg != nil {
		fmt.Println(msg)
	}

}
func (l *LoggerConsole) TRACE(format string, a ...interface{}) {
	if msg := mLog(TRACE, &l.logger, 3, conlog(format, a...)); msg != nil {
		fmt.Println(msg)
	}
}

func (l *LoggerConsole) INFO(format string, a ...interface{}) {
	if msg := mLog(INFO, &l.logger, 3, conlog(format, a...)); msg != nil {
		fmt.Println(msg)
	}
}
func (l *LoggerConsole) WARNING(format string, a ...interface{}) {
	if msg := mLog(WARNING, &l.logger, 3, conlog(format, a...)); msg != nil {
		fmt.Println(msg)
	}
}
func (l *LoggerConsole) ERROR(format string, a ...interface{}) {
	if msg := mLog(ERROR, &l.logger, 3, conlog(format, a...)); msg != nil {
		fmt.Println(msg)
	}
}
func (l *LoggerConsole) FATAL(format string, a ...interface{}) {
	if msg := mLog(FATAL, &l.logger, 3, conlog(format, a...)); msg != nil {
		fmt.Println(msg)
	}
}
