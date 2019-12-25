package errcode

import "encoding/json"

type ErrorCode string

var (
	BlankSymbol    ErrorCode = "Blank symbol"
	InvalidPrice   ErrorCode = "Invalid price"
	EngineExist    ErrorCode = "Engine already exist"
	EngineNotFound ErrorCode = "Engine not exist"
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
		return "交易标的不能为空"
	default:
		return "unknown"
	}
}
