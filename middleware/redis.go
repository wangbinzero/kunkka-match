package middleware

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func Init() {
	addr := viper.GetString("redis.addr")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		//TODO 预留日志输出
	}
}
