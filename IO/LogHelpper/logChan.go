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
	if mes := mLog(DEBUG, &l.logger, 3, conlog(format, a...)); mes != nil {
		l.LogChan <- fmt.Sprintf("%s", mes)
	}

}
func (l *LoggerChan) TRACE(format string, a ...interface{}) {
	if mes := mLog(TRACE, &l.logger, 3, conlog(format, a...)); mes != nil {
		l.LogChan <- fmt.Sprintf("%s", mes)
	}
}

func (l *LoggerChan) INFO(format string, a ...interface{}) {
	if mes := mLog(INFO, &l.logger, 3, conlog(format, a...)); mes != nil {
		l.LogChan <- fmt.Sprintf("%s", mes)
	}

}
func (l *LoggerChan) WARNING(format string, a ...interface{}) {
	if mes := mLog(WARNING, &l.logger, 3, conlog(format, a...)); mes != nil {
		l.LogChan <- fmt.Sprintf("%s", mes)
	}
}
func (l *LoggerChan) ERROR(format string, a ...interface{}) {
	if mes := mLog(ERROR, &l.logger, 3, conlog(format, a...)); mes != nil {
		l.LogChan <- fmt.Sprintf("%s", mes)
	}
}
func (l *LoggerChan) FATAL(format string, a ...interface{}) {
	if mes := mLog(FATAL, &l.logger, 3, conlog(format, a...)); mes != nil {
		l.LogChan <- fmt.Sprintf("%s", mes)
	}
}
