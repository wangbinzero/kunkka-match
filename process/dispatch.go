package process

import (
	"kunkka-match/engine"
	"kunkka-match/enum"
	"kunkka-match/errcode"
	"kunkka-match/middleware/cache"
	"time"
)

func Dispatch(order engine.Order) errcode.ErrorCode {

	//检查交易标的引擎是否开启
	if engine.ChanMap[order.Symbol] == nil {
		return errcode.EngineNotFound
	}

	//如果是创建订单 判断订单缓存是否存在
	//如果存在直接返回错误信息
	if order.Action == enum.ActionCreate {
		if cache.OrderExist(order.Symbol, order.OrderId, order.Action.String()) {
			return errcode.OrderExist
		}
	} else {
		//撤单 订单不存在
		if !cache.OrderExist(order.Symbol, order.OrderId, order.Action.String()) {
			return errcode.OrderNotFound
		} else {
			//撤单，订单存在
		}
	}

	order.Timestamp = time.Now().UnixNano() / 1e3
	cache.SaveOrder(order.ToMap())
	engine.ChanMap[order.Symbol] <- order

	return errcode.OK
}
