package engine

import "kunkka-match/log"

// 保存不同交易标的的订单channel 作为各交易标的的定序队列来用
var ChanMap map[string]chan Order

func Init() {
	ChanMap = make(map[string]chan Order)
	log.Info("order channel [ChanMap] init success\n")
}
