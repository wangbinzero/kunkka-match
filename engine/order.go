package engine

import "kunkka-match/enum"

type Order struct {
	Action  enum.OrderAction `json:"action"`
	Symbol  string           `json:"symbol"`
	OrderId string           `json:"orderId"`
	Side
}
