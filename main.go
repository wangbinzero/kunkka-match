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
