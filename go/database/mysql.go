package database

import (
	"database/sql"
	"fmt"

	"Campus-forum-system/logs"
	"Campus-forum-system/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var sqlDB *sql.DB

func InitMysql() (err error) {
	// 获取配置
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	maxIdleConns := viper.GetInt("mysql.max-idle-conns")
	maxOpenConns := viper.GetInt("mysql.max-open-conns")
	// 格式化
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		username, password, host, port, database)
	// fmt.Println(args)
	// mysql连接
	db, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		// panic("failed to connect mysql database,err:\n" + err.Error())
		logs.Logger.Panicf("连接数据库失败: %s", err.Error())
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns) // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxOpenConns(maxOpenConns) // SetMaxOpenConns 设置打开数据库连接的最大数量。
	} else {
		logs.Logger.Error(err)
	}

	// 自动建表
	if err = db.AutoMigrate(model.GetModelList()...); err != nil {
		logs.Logger.Errorf("自动建表失败: %s", err.Error())
	}

	fmt.Println("数据库初始化完成")
	logs.Logger.Infof("数据库初始化完成")
	return
}

// 返回mysql连接
func GetDB() *gorm.DB {
	return db
}

// 关闭sql连接
func CloseDB() {
	if sqlDB != nil {
		return
	}
	if err := sqlDB.Close(); err != nil {
		logs.Logger.Errorf("关闭数据库连接失败: %s", err.Error())
	}
}
