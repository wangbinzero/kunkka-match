package main

import (
	"kunkka-match/engine"
	"kunkka-match/handler"
	"kunkka-match/log"
	"net/http"
)

// 执行初始化操作
func init() {

	// 1.初始化系统日志
	// 2.引擎初始化
	// 3.中间件初始化
	// 4.
	log.InitLog()

	engine.Init()
	//middleware.Init()
	//process.Init()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/openMatching", handler.OpenMatching)
}
