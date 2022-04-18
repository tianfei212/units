# units 多工具基础模块

1. [时间模块](https://)
2. 文件模块
3. IO基础读写模块
4. 其他模块

## 时间模块:"github.com/tianfei212/units/Sys/DateTimeHelpper"

> 这个模块完成所有的时间操作，包括对时间格式的转换，时间比较，通过时间偏移获得新的时间值

> 1. Format()
>
>> 这个是在模块字段TFormat上附加的函数，这个函数可以完成标准时间格式对golang时间格式的转换
>>
>> ```go
>> fmt.Println(D.TFormat("yyyymmdd").Format()) //20060102
>> ```
>>
>
> 2. 计算时间偏移量时用到的固定预设值
>
>> Min15
>> Min30
>> Min45
>> Hour
>> Hour12
>> Day
>> Week
>>
>
> 3. 统一初始化时支持的预设模式
>
>> StringS// RunTDStr
>> TimeT //RunTDtime
>>
>
> 4. 计算时间属于范围时用到的情况
>
>> ByDAY
>> ByHOUR
>> ByMIN15
>> ByMIN30
>> ByWeek
>>
>
> 5. 初始化函数（RunTDStr） 返回字符串格式的结构体
>
>> ```go
>> d1 := D.RunTDStr("") // 参数为格式化字符串，例如yyyy-mm-dd HH24:mi:ss
>> ```
>>
>
> 6. 初始化函数（RunTDtime） 返回时间格式的结构体
>
>> ```go
>> d2 := D.RunTDtime("yyyy-mm-dd HH24:mi:ss") // 参数为格式化字符串，例如yyyy-mm-dd HH24:mi:ss
>> ```
>>
>
> 7. 初始化函数（RunTD） 返回时间格式的结构体
>
>> ```go
>> d3 := D.RunTD("yyyy-mm-dd HH24:mi:ss", D.TimeT).(D.NewTimer) // 参数为格式化字符串，例如yyyy-mm-dd HH24:mi:ss
>> ```
>>
>
> 8. 方法（Now） 模块的方法，返回字符串格式的时间值,初始化后可以调用
>
>> ```go
>> fmt.Println(d1.Now()) //20220404222341
>>  
>> ```
>>
>
> 9. 方法（NewTime） 模块的方法，返回时间格式的时间值,初始化后可以调用
>
>> 输入参数: SrcTime 输入初始时间，字符串格式
>> 输入参数: IsUTC 布尔值，true，false
>> 输入参数: offset 时间偏移量，支持函数固定值（Min15，Min30，Min45，Hour，Hour12，Day，Week，一个标准字符串或一个数字可以是+-int或+-folat）
>> 输入参数: offsetType 当offset不是固定值时有效，如果offset是一个string则无效
>>
>> ```go
>> d2 := D.RunTDtime("yyyy-mm-dd HH24:mi:ss")
>> fmt.Printf("[%v]--[%v]\n", time.Now(), d2.NewTime(time.Now(), false, D.Min45, "hour"))
>>  
>> ```
>>
>
> 10. 方法（N15M）NewTime的快速方法
> 11. 方法（Min30）NewTime的快速方法
> 12. 方法（N45M）NewTime的快速方法
> 13. 方法（N1Hour）NewTime的快速方法
> 14. 方法（N12Hour）NewTime的快速方法
> 15. 方法（N1Week）NewTime的快速方法
> 16. 方法（CompareDate）计算两个时间值的差（两个时间值得比较，如果stime>dtime (1)),stime=dtime(0),stime<dtime(-1)）
>
>> ```go
>> d1 := D.RunTDStr("")
>> fmt.Println(d1.CompareDate("20201010101010", "20201020101010"))// -1  
>> ```
>> 17. 方法（NumOfDateTime）计算一个时间值是一年中的第几周、日、小时、以及15分钟，30分钟
>>
>
>> ```go
>> d1 := D.RunTDStr("")
>>  fmt.Println(d1.NumOfDateTime("20201010101010", D.ByDAY)) //Day_284
>> ```
>>
## 日志模块:import logh "github.com/tianfei212/units/IO/LogHelpper"
 >这个模块是自写模块，支持打印到屏幕，文件，chan。 
 日志级别支持：
 >```go
> const (
 > UNKNOW LogLevel = iota
 > DEBUG
 > TRACE
 > INFO
 > WARNING
 > ERROR
 > FATAL
 > )
> ```
> 在日志模块初始化时需要指定最大日志等级，将用到上面的部分
### 1)打印到console
>```go
	logc := logh.NewLogConsole("debug")
	logc.DEBUG("a1")
>```
### 2)打印到chan,需要提前设定chan和启动接收函数并启动
> ```go
    lc := make(chan string, 100)
	logchan := logh.NewLogChan("debug", lc)
	go func() {
		//fmt.Println("start show chan log ")
		for v := range lc {
			fmt.Printf("From chan get %v\n", v)
		}
	}()
    logchan.TRACE("this is from chan:%v", i)
>```
### 3)打印到文件，支持文件按照时间，大小，行数进行分隔。
> ```go
>// 初始化控制结构体
> nf := logh.FileSet{
> FilePath:       "o:/TMP/logs", // 设定输出日志的目录，系统会自动创建，如果不存在
> FileName:       "test.log", //日志的基本名称（如果是按照时间切割，时间部分会在文件名前）
> ByModel:        logh.ByFileTime, // ByFileTime/ByFileSize/ByFileRows 
> TimeFormat:     "YYYYMMDDHHMI", // 按照时间切割的时间格式
> MaxFIleSize:    1, // 按照文件大小切割时的大小（MB）
> MaxFIleSaveNum: 10, // 最多保存的文件数量
> MaxFileRows:    1000, // 按照文件行数切割时最大行数
> }
> logf := logh.NewLogFILE("trace", nf)
> logf.StartCheckFile() // 按照文件模式输出时必须在初始化后启动初次建立函数
> logf.INFO("info")
> ```
## 读写配置文件
> 支持读和写 :import rconfig "github.com/tianfei212/units/IO/RwConfig"
### write ()
> ```go
> fmt.Println("写配置文件")
> b := rconfig.ConfigF{FilePath: "o:/tmp/1.json"}
> cc := make(map[string]interface{})
> cc["ccc1"] = 1
> cc["ccc1.bbb"] = "casdt"
> b.Write(cc)
>```
### read()
> ```go
> b := rconfig.ConfigF{FilePath: "o:/tmp/1.json"}
> cs := b.Read("ccc1")
> fmt.Println(cs) //map[ccc1.bbb:casdt]
> ```
## 运行系统命令
```go
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
	b := cmd.OsCommandExec("dir")
	fmt.Println(b)
}

```