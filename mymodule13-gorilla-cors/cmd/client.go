package main

import (
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
	//	rw.Write([]byte("xxxxxxxx TODO: use http's file server to serve a html page with same http request"))
	//})
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, "../webapp/index.html")
	})
	log.Println("starting webapp server on :8282 ... open http://localhost:8282 in the browser!")
	http.ListenAndServe(":8282", nil)
}
