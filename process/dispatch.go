package process

import (
	"kunkka-match/engine"
	"kunkka-match/enum"
	"kunkka-match/errcode"
)

func Dispatch(order engine.Order) errcode.ErrorCode {
	if engine.ChanMap[order.Symbol] == nil {
		return errcode.EngineNotFound
	}

	if order.Action == enum.ActionCreate {

	}
}
