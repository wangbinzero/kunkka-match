package engine

import (
	"github.com/shopspring/decimal"
	"kunkka-match/enum"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	orderQueue := orderQueue{}
	orderQueue.init(enum.SortDirectionAsc)
	order := Order{
		Action:    "create",
		Symbol:    "btcusdt",
		OrderId:   "123456",
		Side:      "buy",
		Type:      "limit",
		Amount:    decimal.NewFromInt(5),
		Price:     decimal.NewFromInt(7000),
		Timestamp: time.Now().UnixNano(),
	}
	order1 := Order{
		Action:    "create",
		Symbol:    "btcusdt",
		OrderId:   "123457",
		Side:      "buy",
		Type:      "limit",
		Amount:    decimal.NewFromInt(4),
		Price:     decimal.NewFromInt(7000),
		Timestamp: time.Now().UnixNano(),
	}

	order2 := Order{
		Action:    "create",
		Symbol:    "btcusdt",
		OrderId:   "123457",
		Side:      "buy",
		Type:      "limit",
		Amount:    decimal.NewFromInt(4),
		Price:     decimal.NewFromInt(7400),
		Timestamp: time.Now().UnixNano(),
	}

	orderQueue.addOrder(order)
	orderQueue.addOrder(order1)
	orderQueue.addOrder(order2)
}
