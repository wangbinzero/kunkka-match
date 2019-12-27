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
	log.Info("Engine [%s] startup success\n", symbol)
	for {
		order, ok := <-ChanMap[symbol]
		if !ok {
			log.Info("Engine [%s] is not running\n", symbol)
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
