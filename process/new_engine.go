package process

import (
	"github.com/shopspring/decimal"
	"kunkka-match/engine"
	"kunkka-match/errcode"
	"kunkka-match/middleware/cache"
)

func NewEngine(symbol string, price decimal.Decimal) errcode.ErrorCode {
	if engine.ChanMap[symbol] != nil {
		return errcode.EngineExist
	}

	//初始化缓冲大小为 100的订单通道
	engine.ChanMap[symbol] = make(chan engine.Order, 100)
	go engine.Run(symbol, price)

	//写入缓存
	cache.SaveSymbol(symbol)
	cache.SavePrice(symbol, price.String())
	return errcode.OK
}
