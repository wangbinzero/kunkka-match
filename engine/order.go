package engine

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"kunkka-match/common"
	"kunkka-match/enum"
	"strconv"
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

// 订单转换为json
func (this Order) toJson() []byte {
	bytes, _ := json.Marshal(&this)
	return bytes
}

//json解析为订单
func (this Order) fromJson(data []byte) {
	json.Unmarshal(data, &this)
}

func (this Order) FromMap(data []interface{}) {
	this.Symbol = data[0].(string)
	this.OrderId = data[1].(string)
	s, _ := strconv.ParseFloat(data[2].(string), 64)
	this.Timestamp = common.Wrap(s, 10)
	this.Action = enum.OrderAction(data[3].(string))
}
