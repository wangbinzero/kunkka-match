package process

import (
	"kunkka-match/engine"
	"kunkka-match/enum"
	"kunkka-match/errcode"
	"kunkka-match/middleware/cache"
	"time"
)

// dispatch to order handler
func Dispatch(order engine.Order) errcode.ErrorCode {
	if engine.ChanMap[order.Symbol] == nil {
		return errcode.EngineNotFound
	}

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
	engine.ChanMap[order.Symbol] <- order
	return errcode.OK
}
