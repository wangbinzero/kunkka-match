package process

import (
	"kunkka-match/engine"
	"kunkka-match/errcode"
)

// 关闭该交易标的引擎
func CloseEngine(symbol string) errcode.ErrorCode {
	if engine.ChanMap[symbol] == nil {
		return errcode.EngineNotFound
	}
	close(engine.ChanMap[symbol])
	return errcode.OK
}
