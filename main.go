package main

import (
	"kunkka-match/common"
	"kunkka-match/conf"
	"kunkka-match/engine"
	"kunkka-match/handler"
	"kunkka-match/log"
	"kunkka-match/middleware"
	"kunkka-match/middleware/mq"
	"kunkka-match/process"
	"net/http"
)

// 执行初始化操作
func init() {
	conf.LoadConfig()
	log.InitLog()
	log.Info(common.Logo)
	engine.Init()
	middleware.Init()
	process.Init()
	//mq.InitAmqp()
	mq.InitEngineMQ()

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/openMatching", handler.OpenMatching)
	mux.HandleFunc("/closeMatching", handler.CloseMatching)
	mux.HandleFunc("/handleOrder", handler.HandleOrder)

	port := conf.Gconfig.GetString("server.port")
	log.Info("服务端口: %s\n", port)
	http.ListenAndServe(port, mux)

}
