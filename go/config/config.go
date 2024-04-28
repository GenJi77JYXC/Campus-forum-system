package config

import (
	"os"

	"github.com/spf13/viper"
)

func ConfigInit() {
	// main函数的工作目录
	workDir, _ := os.Getwd()
	// 配置文件名字
	viper.SetConfigName("config")
	// 配置文件类型
	viper.SetConfigType("yaml")
	// 配置文件路径
	viper.AddConfigPath(workDir + "/config")
	// fmt.Println("配置文件路径" + workDir + "/config")
	// 尝试读入配置文件
	err := viper.ReadInConfig()
	// 读去失败报错退出
	if err != nil {
		panic(err)
	}

}
