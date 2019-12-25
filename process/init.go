package process

import (
	"kunkka-match/engine"
	"kunkka-match/middleware/cache"
)

func Init() {

	//读取已经开启的交易标的
	symbols := cache.GetSymbols()
	for _, symbol := range symbols {

		//根据交易标的查询最新价格
		price := cache.GetPrice(symbol)
		NewEngine(symbol, price)

		//查询缓存中的订单ID列表
		orderIds := cache.GetOrderIdsWithAction(symbol)
		for _, orderId := range orderIds {
			//查询缓存中订单对象
			mapOrder := cache.GetOrder(symbol, orderId)
			order := &engine.Order{}
			order.FromMap(mapOrder)
			engine.ChanMap[order.Symbol] <- *order
		}
	}
}
