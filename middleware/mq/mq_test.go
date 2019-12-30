package mq

import (
	"kunkka-match/middleware"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	middleware.Init()
	SendCancelResult("btcusdt", "123456", true)
	//ConsumerStream("kunkka:match:cancelresults:btcusdt")
	time.Sleep(10 * time.Second)
}
