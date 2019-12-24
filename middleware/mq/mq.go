package mq

import (
	"github.com/go-redis/redis"
	"kunkka-match/common"
	"kunkka-match/middleware"
)

// 发送撤单消息
func SendCancelResult(symbol, orderId string, ok bool) {
	value := map[string]interface{}{"orderId": orderId, "ok": ok}
	a := &redis.XAddArgs{
		Stream:       common.OrderCancelStream + symbol,
		MaxLenApprox: 1000,
		Values:       value,
	}

	middleware.RedisClient.XAdd(a)
}

//发送交易消息
func SendTrade(symbol string, trade map[string]interface{}) {
	a := &redis.XAddArgs{
		Stream:       common.TradeStream + symbol,
		MaxLenApprox: 1000,
		Values:       trade,
	}
	middleware.RedisClient.XAdd(a)
}

//func ConsumerStream(stream string) {
//	fmt.Println("消费 stream: ",stream)
//	for msg := range middleware.RedisClient.XReadStreams(stream).Val() {
//		fmt.Println("哈哈哈哈 stream: ", msg)
//	}
//}
