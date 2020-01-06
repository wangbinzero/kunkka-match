package handler

var (
	SUCCESS string = ""
)

type Response struct {
}

func Success(data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": ""}
}
