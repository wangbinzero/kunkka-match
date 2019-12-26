package main

import (
	"fmt"
	"kunkka-match/common"
	"kunkka-match/middleware"
	"kunkka-match/middleware/mq"
)

func main() {
	middleware.Init()

	mq.SendCancelResult("btcusdt", "12345", true)
	deal()

}

func deal() {

	streams := middleware.RedisClient.XReadStreams(common.OrderCancelStream+"btcusdt", "0").Val()

	for i := 0; i < len(streams); i++ {
		stream := streams[i]
		messages := stream.Messages
		for _, v := range messages {
			fmt.Println("Stream: ", stream)
			fmt.Println("Message: ", v)
			middleware.RedisClient.XAck(stream.Stream, "", v.ID)
		}
	}
}
