package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"kunkka-match/conf"
	"kunkka-match/log"
	"sync"
	"time"
)

var (
	mqConn *amqp.Connection
	mqChan *amqp.Channel
)

//生产者接口
type Producer interface {
	MessageContent() string
}

//消费者接口
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
		log.Info("消息队列连接失败: %v\n", err.Error())
		return
	}

	r.connection = conn
	mqChan, err = conn.Channel()
	r.channel = mqChan
	if err != nil {
		log.Info("消息队列打开channel失败: %v\n", err.Error())
		return
	}
	declareExchange()
}

func (r *RabbitMq) close() {
	err := r.channel.Close()
	if err != nil {
		log.Info("close channel failed: %v\n", err.Error())
		return
	}
	err = r.connection.Close()
	if err != nil {
		log.Info("close amqp connection failed: %v\n", err.Error())
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
			log.Info("declare queue [%s] error: %v\n", r.queue, err.Error())
			return
		}
	}

	err = r.channel.QueueBind(r.queue, r.routeKey, r.exchange, true, nil)
	if err != nil {
		log.Info("queue [%s] bind error: %v\n", r.queue, err.Error())
		return
	}

	//检查交换机是否存在，已经存在不需要重复声明
	err = r.channel.ExchangeDeclarePassive(r.exchange, r.exchangeType, true, false, false, true, nil)
	if err != nil {
		err = r.channel.ExchangeDeclare(r.exchange, r.exchangeType, true, false, false, true, nil)
		if err != nil {
			log.Info("declare exchange [%s] error: %v\n", r.exchange, err.Error())
			return
		}
	}

	//开始发送消息
	err = r.channel.Publish(r.exchange, r.routeKey, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(producer.MessageContent()),
		DeliveryMode: 2,
	})

	if err != nil {
		log.Info("mq publish message failed: %v\n", err.Error())
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
			log.Info("声明队列 [%s] 失败: %v\n", r.queue, err.Error())
			return
		}
	}

	err = r.channel.QueueBind(r.queue, r.routeKey, r.exchange, true, nil)
	if err != nil {
		log.Info("队列 [%s] 绑定失败: %v\n", r.queue, err.Error())
		return
	}
	err = r.channel.Qos(1, 0, true)
	msgList, err := r.channel.Consume(r.queue, "", false, false, false, false, nil)
	if err != nil {
		log.Info("消费消息失败: %v\n", err.Error())
		return
	}
	for {
		msg := <-msgList
		err := receiver.Consumer(msg.Body)
		if err != nil {
			err = msg.Ack(true)
			if err != nil {
				log.Info("消息确认失败: %v\n", err.Error())
				return
			}
		} else {
			err = msg.Ack(false)
			if err != nil {
				log.Info("消息确认失败: %v\n", err.Error())
				return
			}
		}
	}
}

// declare amqp exchange
func declareExchange() {
	matchEx := conf.Gconfig.GetString("rabbitmq.exchange.match.key")
	matchExType := conf.Gconfig.GetString("rabbitmq.exchange.match.type")

	//declare match exchange
	err := mqChan.ExchangeDeclare(matchEx, matchExType,
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Error("声明交换机 [%s] 失败: %v\n", matchEx, err.Error())
		return
	}

	cancelEx := conf.Gconfig.GetString("rabbitmq.exchange.cancel.key")
	cancelExType := conf.Gconfig.GetString("rabbitmq.exchange.cancel.type")

	//declare cancel exchange
	err = mqChan.ExchangeDeclare(cancelEx, cancelExType,
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Error("声明交换机 [%s] 失败: %v\n", cancelEx, err.Error())
		return
	}
	declareQueue()
}

//declare queue
func declareQueue() {
	var err error
	matchQueue := conf.Gconfig.GetString("rabbitmq.queue.match.key")
	cancelQueue := conf.Gconfig.GetString("rabbitmq.queue.cancel.key")
	_, err = mqChan.QueueDeclare(matchQueue, true, false, false, false, nil)
	if err != nil {
		log.Error("声明队列 [%s] 失败: %v\n", matchQueue, err.Error())
		return
	}

	_, err = mqChan.QueueDeclare(cancelQueue, true, false, false, false, nil)
	if err != nil {
		log.Error("声明队列 [%s] 失败: %v\n", cancelQueue, err.Error())
		return
	}

	bindQueue(matchQueue, cancelQueue)
}

func bindQueue(matchQueue, cancelQueue string) {
	mqChan.QueueBind(matchQueue, "match", conf.Gconfig.GetString("rabbitmq.exchange.match.key"), false, nil)
	mqChan.QueueBind(cancelQueue, "cancel", conf.Gconfig.GetString("rabbitmq.exchange.cancel.key"), false, nil)
}
