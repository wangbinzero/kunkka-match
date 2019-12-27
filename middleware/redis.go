package middleware

import (
	"github.com/go-redis/redis"
	"kunkka-match/log"
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
		log.Error("Connection redis error: %v\n", err.Error())
		panic(err)
	} else {
		log.Info("Connection redis[%v] success \n", RedisClient.Options().Addr)
	}
}
