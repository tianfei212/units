package LogHelpper

//////////////////////////
//IO/LogHelpper/LogHelpper.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/25/2022 19:01
// note =
/////////////////////////

type LogLevel uint16

// 定义日志等级
const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// 定义日志输出方向
const (
	O_CONSOLE = 0x00000
	O_FILE    = 0x00010
	O_CHAN    = 0x00100
)

// Logger 定义日志结构体
type logger struct {
	Level LogLevel
	//	IsOutTime   bool
	IsOutRowNum bool
}
type Writer interface {
	DEBUG(a ...interface{})
	TRACE(a ...interface{})
	INFO(a ...interface{})
	WARNING(a ...interface{})
	ERROR(a ...interface{})
	FATAL(a ...interface{})
}
type LoggerConsole struct {
	logger
}
type LoggerFile struct {
	FileSet
	logger
}

const (
	BYFileSize = iota
	BYFileRow
	ByFileTime
)

type FileSet struct {
	FilePath       string
	FileName       string
	SpiltFile      bool
	ByModel        int
	MaxFIleSize    uint64 //byte
	MaxFIleSaveNum int
	TimeFormat     string
	MaxFileRows    uint64
}
type LoggerChan struct {
	LogChan chan string
	logger
}

// NewLogFILE 完成构造函数
func NewLogFILE(maxLogLevel string, F FileSet) *LoggerFile {
	defer func() {
		err := recover()
		if err != nil {
			return
		}
	}()
	level, err := parselogLevel(maxLogLevel)
	if err != nil {
		panic(err)
	}
	lv := logger{
		Level: level,
		//IsOutTime:   true,
		IsOutRowNum: true,
	}
	go writeLogFromChan()
	//createPath(F.FilePath)

	//fmt.Println("cc")
	return &LoggerFile{
		FileSet: F,
		logger:  lv,
	}

}

func NewLogChan(maxLogLevel string, ch chan string) *LoggerChan {
	level, err := parselogLevel(maxLogLevel)
	if err != nil {
		panic(err)
	}
	lv := logger{
		Level: level,

		IsOutRowNum: true,
	}
	return &LoggerChan{
		LogChan: ch,
		logger:  lv,
	}
}

func NewLogConsole(maxLogLevel string) *LoggerConsole {
	level, err := parselogLevel(maxLogLevel)
	if err != nil {
		panic(err)
	}
	return &LoggerConsole{logger{
		Level: level,

		IsOutRowNum: true,
	}}
}
