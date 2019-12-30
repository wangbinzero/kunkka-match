package mq

import (
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"kunkka-match/common"
	"kunkka-match/conf"
	"kunkka-match/log"
	"kunkka-match/middleware"
	"time"
)

var (
	AmqpConnect *amqp.Connection
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

func InitAmqp() {
	var err error
	AmqpConnect, err = amqp.Dial("amqp://" + conf.Gconfig.GetString("rabbitmq.username") + ":" + conf.Gconfig.GetString("rabbitmq.password") + "@" + conf.Gconfig.GetString("rabbitmq.host") + ":" + conf.Gconfig.GetString("rabbitmq.port"+conf.Gconfig.GetString("rabbitmq.vhost")))
	if err != nil {
		log.Info("RabbitMQ connection failed, start reconnect, address: [%s:%s]\n", conf.Gconfig.GetString("rabbitmq.host"), conf.Gconfig.GetString("rabbitmq.port"))
		time.Sleep(5000)
		InitAmqp()
		return
	}

	//if close then reconnect amqp
	go func() {
		<-AmqpConnect.NotifyClose(make(chan *amqp.Error))
		InitAmqp()
	}()

}
