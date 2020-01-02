package handler

import (
	"encoding/json"
	"io/ioutil"
	"kunkka-match/errcode"
	"kunkka-match/process"
	"net/http"
	"strings"
)

// 关闭交易标的
func CloseMatching(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var params closeParams
	if err := json.Unmarshal(body, &params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, v := range params.Symbol {
		if strings.TrimSpace(v) != "" {
			process.CloseEngine(v)
		}
	}
	w.Write(errcode.OK.ToJson())

}

type closeParams struct {
	Symbol []string `json:"symbol"`
}
