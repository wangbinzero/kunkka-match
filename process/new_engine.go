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

	engine.ChanMap[symbol] = make(chan engine.Order, 100)
	go engine.Run(symbol, price)

	cache.SaveSymbol(symbol)
	cache.SavePrice(symbol, price.String())
	return errcode.OK
}
