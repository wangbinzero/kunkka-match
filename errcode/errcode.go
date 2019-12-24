package errcode

import "encoding/json"

type ErrorCode string

var (
	BlankSymbol    ErrorCode = "交易标的不能为空"
	InvalidPrice   ErrorCode = "价格错误"
	EngineExist    ErrorCode = "退出"
	EngineNotFound ErrorCode = "交易标的引擎不存在"
	OK             ErrorCode = ""
)

func (this *ErrorCode) ToJson() []byte {
	bytes, _ := json.Marshal(this)
	return bytes
}
