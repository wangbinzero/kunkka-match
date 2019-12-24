package engine

import (
	"github.com/shopspring/decimal"
	"kunkka-match/enum"
	"kunkka-match/log"
)

//启动撮合引擎
func Run(symbol string, price decimal.Decimal) {
	//lastTradePrice := price
	book := &OrderBook{}

	//初始化订单簿
	book.init()
	log.Info("引擎 [%s] 启动成功", symbol)
	for {
		order, ok := <-ChanMap[symbol]
		if !ok {
			log.Info("引擎 [%s] 未启动", symbol)
			delete(ChanMap, symbol)

			//TODO 清理cache缓存
			return
		}
		log.Info("引擎 [%s] 收到订单: %s", symbol, order.toJson())

		switch order.Action {
		case enum.ActionCreate:
			//TODO 创建订单
		case enum.ActionCancel:
			//TODO 撤销订单
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
