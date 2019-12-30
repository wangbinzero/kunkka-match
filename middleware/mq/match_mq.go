package mq

import "fmt"

type Msg struct {
	Content string
}

func (m *Msg) MessageContent() string {
	return m.Content
}

func (m *Msg) Consumer(data []byte) error {
	fmt.Println("Amqp receive message: ", string(data))
	return nil
}

func InitEngineMQ() {
	msg := fmt.Sprintf("This is test message")

	te := &Msg{Content: msg}

	queueExchange := &QueueExchange{
		QueueName:    "test.rabbit",
		RouteKey:     "rabbit.key",
		ExchangeName: "test.rabbit.mq",
		ExchangeType: "direct",
	}

	mq := NewAmqp(queueExchange)
	mq.RegisterProducer(te)
	mq.RegisterReceiver(te)
	mq.Start()
}
