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
	log.Info("Engine [%s] startup success", symbol)
	for {
		order, ok := <-ChanMap[symbol]
		if !ok {
			log.Info("Engine [%s] is not running", symbol)
			delete(ChanMap, symbol)
			cache.RemoveSymbol(symbol)
			cache.RemovePrice(symbol)
			return
		}
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
		book.removeBuyOrder(order)
	case enum.SideSell:
		book.removeSellOrder(order)
	}

	//TODO 移除缓存

	//TODO 发送到消息队列
	log.Info("Engine: [%s],orderId: [%s] cancelResult: %v\n", order.Symbol, order.OrderId, ok)
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

//限价挂单
func dealLimit(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {
	switch order.Side {
	case enum.SideBuy:
		dealBuyLimit(order, book, lastTradePrice)
	case enum.SideSell:
		dealSellLimit(order, book, lastTradePrice)
	}
}

//限价挂单  -- 买单
func dealBuyLimit(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {
	log.Info("Receive buy limit order: %s", order.toJson())
LOOP:
	headOrder := book.getHeadSellOrder()
	if headOrder == (Order{}) || order.Price.LessThan(headOrder.Price) {
		book.addBuyOrder(*order)
		log.Info("Engine %s, a order has added to the orderBook: %s", order.Symbol, order.toJson())
	} else {
		mathTrade(headOrder, order, book, lastTradePrice)
		if order.Amount.IsPositive() {
			goto LOOP
		}
	}
}

//限价挂单 -- 卖单
func dealSellLimit(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {
	log.Info("Receive sell limit order: %s", order.toJson())
}

//成交撮合
//做数量减法即可
func mathTrade(headOrder Order, order *Order, book *OrderBook, lastTradePrice decimal.Decimal) *Order {
	result := order.Amount.Sub(headOrder.Amount)
	order.Amount = result
	if result.Cmp(decimal.Zero) > 0 {

	}
	return order
}
