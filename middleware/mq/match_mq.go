package mq

import (
	"encoding/json"
	"kunkka-match/engine"
	"kunkka-match/errcode"
	"kunkka-match/log"
	"kunkka-match/process"
)

type Msg struct {
	Content string
}

func (m *Msg) MessageContent() string {
	return m.Content
}

// 消费消息队列
func (m *Msg) Consumer(data []byte) error {

	var order engine.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		return err
	}
	errco := process.Dispatch(order)
	if errco.String() != errcode.OK.String() {

		//TODO 订单已存在于订单簿中，是否需要返回调用端数据已存在？
		log.Error("消费订单消息失败: 订单号: [%s] %s\n", order.OrderId, errco.String())
	}
	return nil
}

func InitEngineMQ() {
	te := &Msg{}

	queueExchange := &QueueExchange{
		QueueName:    "kunkka.queue.match",
		RouteKey:     "match",
		ExchangeName: "kunkka.exchange.match",
		ExchangeType: "direct",
	}
	mq := NewAmqp(queueExchange)
	mq.RegisterProducer(te)
	mq.RegisterReceiver(te)
	mq.Start()
}
