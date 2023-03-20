package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	devCfg = "./config.yaml"
)

// Init 读取配置文件
// 会依次到 linux默认路径 > windows默认路径 > local 寻找配置文件，全部没找到时候panic
func Init() {
	cfgFile := devCfg
	if _, err := os.Stat(cfgFile); err == nil {
		initViper(cfgFile)
		return
	}
}

func initViper(file string) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Read config error: " + err.Error())
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
