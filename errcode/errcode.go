package errcode

import "encoding/json"

type ErrorCode string

var (
	BlankSymbol  ErrorCode = "交易标的不能为空"
	InvalidPrice ErrorCode = "价格错误"
	EngineExist  ErrorCode = "退出"
)

func (this *ErrorCode) ToJson() []byte {
	bytes, _ := json.Marshal(this)
	return bytes
}
