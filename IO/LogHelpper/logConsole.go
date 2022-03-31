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
	fmt.Println(mLog(DEBUG, &l.logger, 3, conlog(format, a...)))
}
func (l *LoggerConsole) TRACE(format string, a ...interface{}) {
	fmt.Println(mLog(TRACE, &l.logger, 3, conlog(format, a...)))
}

func (l *LoggerConsole) INFO(format string, a ...interface{}) {
	fmt.Println(smalllog(INFO, &l.logger, format, a...))

}
func (l *LoggerConsole) WARNING(format string, a ...interface{}) {
	fmt.Println(smalllog(WARNING, &l.logger, format, a...))
}
func (l *LoggerConsole) ERROR(format string, a ...interface{}) {
	fmt.Println(smalllog(ERROR, &l.logger, format, a...))
}
func (l *LoggerConsole) FATAL(format string, a ...interface{}) {
	fmt.Println(smalllog(FATAL, &l.logger, format, a...))
}
