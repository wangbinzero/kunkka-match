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
	book.init()
	log.Info("撮合引擎 [%s] 启动成功\n", symbol)
	for {
		order, ok := <-ChanMap[symbol]
		if !ok {
			log.Info("撮合引擎 [%s] 未开启\n", symbol)
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
