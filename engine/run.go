package engine

import (
	"github.com/shopspring/decimal"
	"kunkka-match/enum"
	"kunkka-match/log"
	"kunkka-match/middleware/cache"
)

//启动撮合引擎
func Run(symbol string, price decimal.Decimal) {
	lastTradePrice := price
	book := &OrderBook{}

	//初始化订单簿
	book.init()
	log.Info("引擎 [%s] 启动成功", symbol)
	for {
		order, ok := <-ChanMap[symbol]
		if !ok {
			log.Info("引擎 [%s] 未启动", symbol)
			delete(ChanMap, symbol)
			cache.RemoveSymbol(symbol)
			cache.RemovePrice(symbol)
			return
		}
		log.Info("引擎 [%s] 收到订单: %s", symbol, order.toJson())

		switch order.Action {
		case enum.ActionCreate:

			dealCreate(&order, book, lastTradePrice)
		case enum.ActionCancel:

			dealCancel(&order, book)
		}
	}

}

// 撤单
func dealCancel(order *Order, book *OrderBook) {
	var ok bool
	switch order.Side {
	case enum.SideBuy:
		ok = book.removeBuyOrder(order)
	case enum.SideSell:
		ok = book.removeSellOrder(order)
	}

	//TODO 移除缓存

	//TODO 发送到消息队列
	log.Info("引擎 [%s],订单 [%s] 撤单结果: %s", order.Symbol, order.OrderId, ok)
}

// 创建订单
func dealCreate(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {
	switch order.Type {
	case enum.Limit:
		dealLimit(order, book, lastTradePrice)
	case enum.LimitIoc:
		//dealLimitIoc(order, book, lastTradePrice)
	case enum.Market:
		//dealMarket(order, book, lastTradePrice)
	case enum.MarketTop5:
		//dealMarketTop5(order, book, lastTradePrice)
	case enum.MarketTop10:
		//dealMarketTop10(order, book, lastTradePrice)
	case enum.MarketOpponent:
		//dealMarketOpponent(order, book, lastTradePrice)

	}
}

func dealLimit(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {
	switch order.Side {
	case enum.SideBuy:
	case enum.SideSell:

	}
}
