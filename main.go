package main

import (
	"kunkka-match/conf"
	"kunkka-match/engine"
	"kunkka-match/handler"
	"kunkka-match/log"
	"kunkka-match/middleware"
	"kunkka-match/process"
	"net/http"
)

// 执行初始化操作
func init() {

	// 1.初始化系统日志
	// 2.引擎初始化
	// 3.中间件初始化
	// 4.
	conf.LoadConfig()
	log.InitLog()

	engine.Init()
	middleware.Init()
	process.Init()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/openMatching", handler.OpenMatching)
	mux.HandleFunc("/closeMatching", handler.CloseMatching)
	mux.HandleFunc("/handleOrder", handler.HandleOrder)
	//log.Info("")
}
