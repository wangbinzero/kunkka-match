package main

import (
	"kunkka-match/engine"
	"kunkka-match/handler"
	"kunkka-match/log"
	"kunkka-match/middleware"
	"net/http"
)

// 执行初始化操作
func init() {
	log.InitLog()

	engine.Init()
	middleware.Init()
	process.Init()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/openMatching", handler.OpenMatching)
}
