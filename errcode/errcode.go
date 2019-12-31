package errcode

import "encoding/json"

type ErrorCode string

var (
	BlankSymbol    ErrorCode = "blank symbol"
	InvalidPrice   ErrorCode = "invalid price"
	EngineExist    ErrorCode = "engine already exist"
	EngineNotFound ErrorCode = "engine not found"
	OrderExist     ErrorCode = "order exist"
	OrderNotFound  ErrorCode = "order not found"
	OK             ErrorCode = "ok"
)

func (this ErrorCode) ToJson() []byte {
	bytes, _ := json.Marshal(&this)
	return bytes
}

func (this ErrorCode) IsOk() bool {
	if this.String() == "ok" {
		return true
	}
	return false
}

func (this ErrorCode) String() string {
	switch this {
	case BlankSymbol:
		return "blank symbol"
	case InvalidPrice:
		return "invalid price"
	case EngineExist:
		return "engine already exist"
	case EngineNotFound:
		return "engine not found"
	case OrderExist:
		return "order exist"
	case OrderNotFound:
		return "order not found"
	case OK:
		return "ok"
	default:
		return "unknown"
	}
}
