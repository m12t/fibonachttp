package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const host = "http://127.0.0.1:8080"

type payload struct {
	Num int `json:"ans"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fibonacci/", fibonacci)
	http.ListenAndServe("127.0.0.1:8080", mux)
}

func fibonacci(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if n <= 1 {
		ans, _ := json.Marshal(payload{Num: n})
		w.Write(ans)
		return
	}
	var p1, p2 payload
	r1, _ := http.Get(fmt.Sprintf("%s/fibonacci/%d", host, n-1))
	json.NewDecoder(r1.Body).Decode(&p1)
	r2, _ := http.Get(fmt.Sprintf("%s/fibonacci/%d", host, n-2))
	json.NewDecoder(r2.Body).Decode(&p2)
	ans, _ := json.Marshal(payload{Num: p1.Num + p2.Num})
	w.Write(ans)
}
