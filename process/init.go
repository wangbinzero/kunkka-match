package process

import (
	"kunkka-match/engine"
	"kunkka-match/middleware/cache"
)

func Init() {

	symbols := cache.GetSymbols()
	for _, symbol := range symbols {

		price := cache.GetPrice(symbol)
		NewEngine(symbol, price)

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
