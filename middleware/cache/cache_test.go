package cache

import (
	"kunkka-match/middleware"
	"testing"
)

// 保存交易标的
func TestSaveSymbol(t *testing.T) {
	middleware.Init()
	SaveSymbol("btcusdt")
	SaveSymbol("eosusdt")

}

// 查询交易标的列表
func TestGetSymbols(t *testing.T) {
	middleware.Init()
	GetSymbols()
}

// 移除交易标的
func TestRemoveSymbol(t *testing.T) {
	middleware.Init()
	RemoveSymbol("btcusdt")
}

func TestSavePrice(t *testing.T) {
	middleware.Init()
	SavePrice("btcusdt", "7600")
}

func TestGetPrice(t *testing.T) {
	middleware.Init()
	GetPrice("btcusdt")
}
