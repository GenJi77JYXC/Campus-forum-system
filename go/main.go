package main

import (
	"Campus-forum-system/config"
	"Campus-forum-system/database"
	"Campus-forum-system/logs"
	"Campus-forum-system/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.ConfigInit()
	logs.InitLogger("./logs/log", 1, 1, 2, false) // 在./logs/log目录下创建日志文件，日志文件最大为1M，保留最近3个日志文件，日志级别为debug，不输出到控制台
	database.InitMysql()
	defer database.CloseDB()

	r := gin.Default()
	r = routers.CollectRouter(r)
	// 从viper中获取到运行端口
	port := viper.GetString("server.port")
	// 如果指定了端口
	if port != "" {
		panic(r.Run(":" + port))
	}
	//	没指定端口就直接运行
	panic(r.Run())
}
