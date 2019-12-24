package engine

import (
	"github.com/shopspring/decimal"
	"kunkka-match/enum"
)

type Order struct {
	Action    enum.OrderAction `json:"action"`
	Symbol    string           `json:"symbol"`
	OrderId   string           `json:"orderId"`
	Side      enum.OrderSide   `json:"side"`
	Type      enum.OrderType   `json:"orderType"`
	Amount    decimal.Decimal  `json:"amount"`
	Price     decimal.Decimal  `json:"price"`
	Timestamp int64            `json:"timestamp"`
}
