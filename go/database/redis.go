package database

import (
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
