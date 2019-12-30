package mq

import (
	"fmt"
	"testing"
)

func TestMsg_Consumer(t *testing.T) {
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
