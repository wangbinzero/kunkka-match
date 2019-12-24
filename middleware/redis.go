package middleware

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init() {
	//addr := conf.Gconfig.GetString("redis.addr")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Redis初始化成功，地址: [%s]", "127.0.0.1:6379")
		//log.Info("Redis初始化成功，地址: [%s]","127.0.0.1:6379")
	}
}
