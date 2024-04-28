package database

import (
	"fmt"

	"Campus-forum-system/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	// 获取配置
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	// 格式化
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		username, password, host, port, database)
	// fmt.Println(args)
	// mysql连接
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql database,err:\n" + err.Error())
	}
	// 自动建表
	db.AutoMigrate(&model.User{})

	DB = db
	fmt.Println("数据库初始化完成")
	return db
}

// 返回mysql连接
func GetDB() *gorm.DB {
	return DB
}
