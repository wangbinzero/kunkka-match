package mq

import (
	"encoding/json"
	"kunkka-match/engine"
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
	process.Dispatch(order)
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
	te.Content = "Hello World"
	mq := NewAmqp(queueExchange)
	mq.RegisterProducer(te)
	mq.RegisterReceiver(te)
	mq.Start()
}
