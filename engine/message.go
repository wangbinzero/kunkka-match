package engine

import (
	"github.com/shopspring/decimal"
	"kunkka-match/log"
)

// 发送成交消息
func SendMessage(order Order, message Trade) {
	if order.Amount.IsPositive() {
		message.Remain = order.Amount
	} else {
		message.Remain = decimal.Zero
	}
	log.Info("最新成交消息: 交易标的 [%s] 交易类型: [%s]  消息: %v\n", order.Symbol, order.OrderType.String(), message.toJson())
}
