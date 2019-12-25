package cache

import (
	"kunkka-match/common"
	"kunkka-match/middleware"
	"testing"
	"time"
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

func TestSaveOrder(t *testing.T) {
	middleware.Init()
	order := make(map[string]interface{})
	order["symbol"] = "btcusdt"
	order["orderId"] = "123456"
	order["timestamp"] = common.Unwrap(time.Now().UnixNano(), 10)
	order["action"] = "buy"
	SaveOrder(order)
}

func TestGetOrder(t *testing.T) {
	middleware.Init()
	GetOrder("btcusdt", "123456")
}
