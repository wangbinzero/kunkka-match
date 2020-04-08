package engine

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"kunkka-match/common"
	"kunkka-match/enum"
	"strconv"
)

/**
{
	"action": "create",
	"symbol": "btcusdt",
	"orderId": "123456",
	"side": "buy",
	"orderType": "limit",
	"amount": "1",
	"price": 8100
}
*/
type Order struct {
	Action    enum.OrderAction `json:"action"`    //创建订单  | 取消订单
	Symbol    string           `json:"symbol"`    //标的
	OrderId   string           `json:"orderId"`   //订单号
	Side      enum.OrderSide   `json:"side"`      //买 | 卖
	OrderType enum.OrderType   `json:"orderType"` //订单类型 限价 | 市价  ...
	Amount    decimal.Decimal  `json:"amount"`    //数量
	Price     decimal.Decimal  `json:"price"`     //价格
	Timestamp int64            `json:"timestamp"` //时间戳
}

// 订单转换为json
func (this Order) toJson() string {
	bytes, _ := json.Marshal(&this)
	return string(bytes)
}

//json解析为订单
func (this Order) fromJson(data []byte) {
	json.Unmarshal(data, &this)
}

//map转对象
func (this *Order) FromMap(data map[string]string) {
	this.Symbol = data["symbol"]
	this.OrderId = data["orderId"]
	s, _ := strconv.ParseFloat(data["timestamp"], 64)
	this.Timestamp = common.Wrap(s, 10)
	this.Action = enum.OrderAction(data["action"])
	this.OrderType = enum.OrderType(data["orderType"])
	this.Side = enum.OrderSide(data["side"])
}

//对象转Map
func (this *Order) ToMap() map[string]interface{} {
	var orderMap = make(map[string]interface{})
	orderMap["symbol"] = this.Symbol
	orderMap["orderId"] = this.OrderId
	orderMap["timestamp"] = common.Unwrap(this.Timestamp, 64)
	orderMap["action"] = this.Action.String()
	orderMap["price"] = this.Price.String()
	orderMap["orderType"] = this.OrderType.String()
	orderMap["side"] = this.Side.String()
	return orderMap
}
