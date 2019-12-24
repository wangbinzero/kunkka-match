package process

import (
	"github.com/shopspring/decimal"
	"kunkka-match/engine"
	"kunkka-match/errcode"
)

func NewEngine(symbol string, price decimal.Decimal) errcode.ErrorCode {
	if engine.ChanMap[symbol] != nil {
		return errcode.EngineExist
	}

	//初始化缓冲大小为 100的订单通道
	engine.ChanMap[symbol] = make(chan engine.Order, 100)
	go engine.Run(symbol, price)
	//TODO
	//cache.SaveSymbol(symbol)
	//cache.SavePrice(symbol,price)
	return errcode.OK
}
