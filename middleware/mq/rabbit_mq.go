package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"kunkka-match/log"
	"sync"
	"time"
)

var (
	mqConn *amqp.Connection
	mqChan *amqp.Channel
)

type Producer interface {
	MessageContent() string
}

type Receiver interface {
	Consumer([]byte) error
}

type RabbitMq struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	queue        string
	routeKey     string
	exchange     string
	exchangeType string
	producerList []Producer
	receiverList []Receiver
	mu           sync.RWMutex
}

type QueueExchange struct {
	QueueName    string //队列名
	RouteKey     string //路由key
	ExchangeName string //交换机名称
	ExchangeType string //交换机类型
}

func (r *RabbitMq) connect() {
	var err error

	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", "guest", "guest", "127.0.0.1", 5672)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Error("connect to amqp failed: %v\n", err.Error())
		return
	}

	r.connection = conn
	mqChan, err = conn.Channel()
	if err != nil {
		log.Error("open channel failed: %v\n", err.Error())
		return
	}
	r.channel = mqChan

}

func (r *RabbitMq) close() {
	err := r.channel.Close()
	if err != nil {
		log.Error("close channel failed: %v\n", err.Error())
		return
	}
	err = r.connection.Close()
	if err != nil {
		log.Error("close amqp connection failed: %v\n", err.Error())
		return
	}
}

func NewAmqp(e *QueueExchange) *RabbitMq {
	return &RabbitMq{
		queue:        e.QueueName,
		routeKey:     e.RouteKey,
		exchange:     e.ExchangeName,
		exchangeType: e.ExchangeType,
	}
}

func (r *RabbitMq) Start() {

	for _, producer := range r.producerList {
		go r.listenProducer(producer)
	}

	for _, receiver := range r.receiverList {
		go r.listenReceiver(receiver)
	}
	time.Sleep(1 * time.Second)
}

func (r *RabbitMq) RegisterProducer(producer Producer) {
	r.producerList = append(r.producerList, producer)
}

func (r *RabbitMq) RegisterReceiver(receiver Receiver) {
	r.mu.Lock()
	r.receiverList = append(r.receiverList, receiver)
	r.mu.Unlock()
}

func (r *RabbitMq) listenProducer(producer Producer) {

	if r.channel == nil {
		r.connect()
	}

	//检查队列是否存在，已经存在不需要重复声明
	_, err := r.channel.QueueDeclarePassive(r.queue, true, false, false, true, nil)
	if err != nil {
		//队列不存在，声明
		_, err = r.channel.QueueDeclare(r.queue, true, false, false, true, nil)
		if err != nil {
			log.Error("declare queue [%s] error: %v\n", r.queue, err.Error())
			return
		}
	}

	err = r.channel.QueueBind(r.queue, r.routeKey, r.exchange, true, nil)
	if err != nil {
		log.Error("queue [%s] bind error: %v\n", r.queue, err.Error())
		return
	}

	//检查交换机是否存在，已经存在不需要重复声明
	err = r.channel.ExchangeDeclarePassive(r.exchange, r.exchangeType, true, false, false, true, nil)
	if err != nil {
		err = r.channel.ExchangeDeclare(r.exchange, r.exchangeType, true, false, false, true, nil)
		if err != nil {
			log.Error("declare exchange [%s] error: %v\n", r.exchange, err.Error())
			return
		}
	}

	//开始发送消息
	err = r.channel.Publish(r.exchange, r.routeKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(producer.MessageContent()),
	})

	if err != nil {
		log.Error("mq publish message failed: %v\n", err.Error())
		return
	}

}

func (r *RabbitMq) listenReceiver(receiver Receiver) {
	defer r.close()

	if r.channel == nil {
		r.connect()
	}

	//检查队列
	_, err := r.channel.QueueDeclarePassive(r.queue, true, false, false, true, nil)
	if err != nil {
		_, err = r.channel.QueueDeclare(r.queue, true, false, false, true, nil)
		if err != nil {
			log.Error("declare queue [%s] error: %v\n", r.queue, err.Error())
			return
		}
	}

	err = r.channel.QueueBind(r.queue, r.routeKey, r.exchange, true, nil)
	if err != nil {
		log.Error("queue [%s] bind error: %v\n", r.queue, err.Error())
		return
	}
	err = r.channel.Qos(1, 0, true)
	msgList, err := r.channel.Consume(r.queue, "", false, false, false, false, nil)
	if err != nil {
		log.Error("receive message error: %v\n", err.Error())
		return
	}

	for msg := range msgList {
		err := receiver.Consumer(msg.Body)
		if err != nil {
			err = msg.Ack(true)
			if err != nil {
				log.Error("confirm message error: %v\n", err.Error())
				return
			}
		} else {
			err = msg.Ack(false)
			if err != nil {
				log.Error("confirm message error: %v\n", err.Error())
				return
			}
			return
		}
	}
}
