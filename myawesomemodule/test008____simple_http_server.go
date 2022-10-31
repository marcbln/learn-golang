package main

import (
	"fmt"
	"net/http"
)

func myFunc(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from http server äöüßα"))
}

func main() {
	fmt.Printf("starting server...\n")
	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":8080", nil)
}
