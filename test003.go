package main

import (
	"net/http"
)

func handleHttp(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello from handleHttp()"))
}

func main() {
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
