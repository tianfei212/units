package LogHelpper

import "fmt"

//////////////////////////
//IO/LogHelpper/logChan.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/31/2022 13:18
// note =
/////////////////////////

func (l *LoggerChan) DEBUG(format string, a ...interface{}) {
	l.LogChan <- fmt.Sprintf("%s", mLog(DEBUG, &l.logger, 3, conlog(format, a...)))

}
func (l *LoggerChan) TRACE(format string, a ...interface{}) {
	l.LogChan <- fmt.Sprintf("%s", mLog(TRACE, &l.logger, 3, conlog(format, a...)))
}

func (l *LoggerChan) INFO(format string, a ...interface{}) {
	//if _, isClose := <-l.LogChan; isClose {
	//	fmt.Println("close chan")
	//} else {
	l.LogChan <- fmt.Sprintf("%s", mLog(INFO, &l.logger, 3, conlog(format, a...)))
	//}

}
func (l *LoggerChan) WARNING(format string, a ...interface{}) {
	l.LogChan <- fmt.Sprintf("%s", mLog(WARNING, &l.logger, 3, conlog(format, a...)))
}
func (l *LoggerChan) ERROR(format string, a ...interface{}) {
	l.LogChan <- fmt.Sprintf("%s", mLog(ERROR, &l.logger, 3, conlog(format, a...)))
}
func (l *LoggerChan) FATAL(format string, a ...interface{}) {
	l.LogChan <- fmt.Sprintf("%s", mLog(FATAL, &l.logger, 3, conlog(format, a...)))
}
