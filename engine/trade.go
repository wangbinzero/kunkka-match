package engine

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"kunkka-match/enum"
	"kunkka-match/log"
	"kunkka-match/middleware/cache"
	"time"
)

type Trade struct {
	MarkerId  string          `json:"markerId"`  //挂单id
	TakerId   string          `json:"takerId"`   //吃单id
	TakerSide string          `json:"takerSide"` //吃单方向
	Amount    decimal.Decimal `json:"amount"`    //成交数量
	Price     decimal.Decimal `json:"price"`     //成交价格
	Timestamp int64           `json:"timestamp"` //成交时间戳
}

//成交对象结果转换为json字符串
func (this Trade) toJson() string {
	bytes, _ := json.Marshal(&this)
	return string(bytes)
}

//成交撮合
//做数量减法即可
func matchTrade(headOrder *Order, order *Order, book *OrderBook, lastTradePrice decimal.Decimal) *Order {
	//上一笔最新成交价
	pPrice := cache.GetPrice(order.Symbol)
	var buyPrice decimal.Decimal
	var sellPrice decimal.Decimal

	if order.Side == enum.SideBuy {
		buyPrice, sellPrice = order.Price, headOrder.Price
	} else {
		buyPrice, sellPrice = headOrder.Price, order.Price
	}

	//计算最新价
	currenDealPrice := newDealPrice(pPrice, buyPrice, sellPrice)
	result := order.Amount.Sub(headOrder.Amount)
	order.Amount = result
	// result > 0 表示订单部分成交 对手单完全成交
	// result = 0 表示订单完全成交 对手单完全成交
	// result < 0 表示订单完全成交 对手单部分成交
	if result.Cmp(decimal.Zero) == 0 {
		switch headOrder.Side {
		case enum.SideBuy:
			book.removeBuyOrder(headOrder)
		case enum.SideSell:
			book.removeSellOrder(headOrder)
		}
		cache.RemoveOrder(headOrder.Symbol, headOrder.OrderId, headOrder.Action.String())
		log.Info("订单完全成交, 吃单id: [%s] 挂单id: [%s]  订单类型: [%v] 成交数量: %v\n", order.OrderId, headOrder.OrderId, order.OrderType, headOrder.Amount)
	} else if result.Cmp(decimal.Zero) > 0 {
		switch order.Side {
		case enum.SideSell:
			book.addSellOrder(*order)
			break
		case enum.SideBuy:
			book.addBuyOrder(*order)
			break
		}

		cache.SaveOrder(order.ToMap())
		cache.SavePrice(order.Symbol, currenDealPrice.String())
		//TODO 交易标的最新价是在此处存如缓存吗
	} else if result.Cmp(decimal.Zero) < 0 {

	}

	//TODO 发送成交数据给行情系统
	// 发送订单成交信息给业务系统
	return order
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
	log.Info("撮合引擎: [%s],订单ID: [%s] 撤单结果: %v\n", order.Symbol, order.OrderId, ok)
}

//撤单逻辑处理
func cancelOrder(order *Order, book *OrderBook) {

}

// 创建订单
func dealCreate(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {
	switch order.OrderType {
	case enum.Limit:
		dealLimit(order, book, lastTradePrice)
	case enum.LimitIoc:
		dealLimitIoc(order, book, lastTradePrice)
	case enum.Market:
		dealMarket(order, book, lastTradePrice)
	case enum.MarketTop5:
		//dealMarketTop5(order, book, lastTradePrice)
	case enum.MarketTop10:
		//dealMarketTop10(order, book, lastTradePrice)
	case enum.MarketOpponent:
		//dealMarketOpponent(order, book, lastTradePrice)

	}
}

//市价，及时成交，剩余撤销
func dealMarket(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {

	switch order.Side {
	case enum.SideBuy:
		dealBuyMarket(order, book, lastTradePrice)
	case enum.SideSell:
		dealSellMarket(order, book, lastTradePrice)
	}
}

func dealBuyMarket(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {

}

func dealSellMarket(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {

}

//订单如果不能立即成交则未成交部分直接撤单
func dealLimitIoc(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {

}

//限价单
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
LOOP:
	headOrder := book.getHeadSellOrder()

	//如果限价买单为空 或者 价格小于卖单，则加入买单簿，不进行撮合
	//否则进行撮合逻辑处理
	if headOrder == (Order{}) || order.Price.LessThan(headOrder.Price) {
		cache.SaveOrder(order.ToMap())
		book.addBuyOrder(*order)
		log.Info("撮合引擎 %s, 添加订单簿数据,买单: %s\n", order.Symbol, order.toJson())
	} else {
		matchTrade(&headOrder, order, book, lastTradePrice)
		var trade Trade
		trade.MarkerId = headOrder.OrderId
		trade.TakerId = order.OrderId
		trade.TakerSide = order.Side.String()
		trade.Timestamp = time.Now().UnixNano() / 1e3

		if order.Amount.IsPositive() {
			goto LOOP
		}
	}
}

//限价挂单 -- 卖单
func dealSellLimit(order *Order, book *OrderBook, lastTradePrice decimal.Decimal) {
	log.Info("限价卖单: %s\n", order.toJson())
LOOP:
	headOrder := book.getHeadBuyOrder()
	if headOrder == (Order{}) || order.Price.GreaterThan(headOrder.Price) {
		book.addSellOrder(*order)
		log.Info("撮合引擎 %s, 添加订单簿数据,卖单: %s\n", order.Symbol, order.toJson())
	} else {
		matchTrade(&headOrder, order, book, lastTradePrice)
		if order.Amount.IsPositive() {
			goto LOOP
		}
	}
}

//计算最新成交价
func newDealPrice(prevDealPrice, buyPrice, sellPrice decimal.Decimal) decimal.Decimal {
	if prevDealPrice.GreaterThanOrEqual(buyPrice) {
		return buyPrice
	} else if prevDealPrice.LessThanOrEqual(sellPrice) {
		return sellPrice
	} else if buyPrice.GreaterThan(prevDealPrice) && prevDealPrice.GreaterThan(sellPrice) {
		return prevDealPrice
	}
	return prevDealPrice
}
