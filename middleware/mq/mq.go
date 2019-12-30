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
		log.Info("rabbitmq connection failed, start reconnect, address: [%s:%s]\n", conf.Gconfig.GetString("rabbitmq.host"), conf.Gconfig.GetString("rabbitmq.port"))
		time.Sleep(5000)
		InitAmqp()
		return
	}

	//if close then reconnect amqp
	go func() {
		<-AmqpConnect.NotifyClose(make(chan *amqp.Error))
		InitAmqp()
	}()

	//declareExchange()
}

// pubish message to amqp server
// deliveryMode =1  non persistent  2 persistent
//func PublishMessage(exchange, routeKey, contentType string, message *[]byte, deliveryMode uint8) {
//	channel, err := AmqpConnect.Channel()
//	defer channel.Close()
//	if err != nil {
//		log.Error("can't get a channel from connection for publish message: %v\n", err.Error())
//		return
//	}
//	if err = channel.Publish(
//		exchange,
//		routeKey,
//		false,
//		false,
//		amqp.Publishing{
//			Headers:         amqp.Table{},
//			ContentType:     contentType,
//			ContentEncoding: "",
//			DeliveryMode:    deliveryMode,
//			Priority:        0,
//			Body:            *message,
//		},
//	); err != nil {
//		log.Error("publish message failed: %v\n", err.Error())
//		return
//	}
//
//	log.Info("publish message success\n")
//}
