package RwConfig

import "C"
import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tianfei212/units/IO/FileHelpper"
	"strings"
)

//////////////////////////
//IO/RwConfig/RwConfig.go
//author = "Derek Tian"
//Ver = 0.0.0.1
//make time = 3/25/2022 15:01
// note = 配置文件读写
/////////////////////////

// Configer 定义配置文件接口，实现读取，写入，标签查找
type Configer interface {
	Read(Key string) map[string]interface{}
	Write(Key, Value string) bool
}

// ConfigF 定义结构体
type ConfigF struct {
	FilePath string
}

// 读的方法，需要传入Key（A.B、A）的结构
func (Cf ConfigF) Read(Key string) map[string]interface{} {
	var mConfig = make(map[string]interface{}, 2)
	//fmt.Println(Cf.FilePath)
	//fmt.Println(Key)
	if FileHelpper.RextStr(Cf.FilePath) == "" {
		fmt.Println("not file Config FIle!!!")
		return mConfig
	} else {

		viper.SetConfigFile(Cf.FilePath)
		//config.AddConfigPath(".")
		//viper.SetConfigType("ini")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		} else {

			if strings.Contains(Key, ".") {
				mConfig[Key] = viper.Get(Key)
			} else {
				for _, v := range viper.AllKeys() {
					if strings.Contains(v, Key) {
						mConfig[v] = viper.Get(v)
					}

				}
			}
		}

		//	fmt.Println(viper.GetString(Key))
		return mConfig

	}

}

// 写的方法
func (Cf ConfigF) Write(m map[string]string) bool {
	if FileHelpper.RextStr(Cf.FilePath) != "" {
		viper.SetConfigFile(Cf.FilePath)
		for k, v := range m {
			viper.Set(k, v)
		}
		err := viper.WriteConfigAs(Cf.FilePath)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
