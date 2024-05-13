package database

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func InitRedis() (err error) {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0, // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {
	return rdb
}

func RedisSetKey(key, val string) {
	// 将该条token失效的时间设置为token失效的最大时间 24*time.Hour*time.Duration(model.TokenExpireDays)
	err := rdb.Set(key, val, 24*time.Hour*time.Duration(model.TokenExpireDays)).Err()
	if err != nil {
		logs.Logger.Errorf("RedisSetKey出错:%s", err)
		fmt.Println("RedisSetKey出错")
		panic(err)
	}
	// defer func(rdb *redis.Client) {
	// 	err := rdb.Close()
	// 	if err != nil {
	// 		fmt.Println("rdb关闭错误：", err)
	// 	}
	// }(rdb)
}

func RedisGetKey(key string) string {
	val, _ := rdb.Get(key).Result()
	// if err != nil {
	// fmt.Println("RedisGetKey出错:", err)
	// fmt.Println("Redis里没有该Key,返回的val值是:", val)
	// }
	// defer func(rdb *redis.Client) {
	// 	err := rdb.Close()
	// 	if err != nil {
	// 		fmt.Println("rdb关闭错误:", err)
	// 	}
	// }(rdb)

	// 如果val等于空字符串， 证明这个key没有被放入redis， 也就是这个token没被注销
	return val
}
