package main

import (
	"kunkka-match/handler"
	"net/http"
)

// 执行初始化操作
func init() {

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/openMatching", handler.OpenMatching)
}
