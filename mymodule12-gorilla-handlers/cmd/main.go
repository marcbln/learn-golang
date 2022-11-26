package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func getRoot(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello from getRoot()\n"))
	rw.WriteHeader(http.StatusOK)
}
func getHello(rw http.ResponseWriter, req *http.Request) {
	if !(req.Method == "PUT" || req.Method == "POST" || req.Method == "PATCH") {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rw.Write([]byte("Hello from getHello()\n"))
	rw.WriteHeader(http.StatusCreated)
}
func getBigResponse(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello from getBigResponse()\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.Handle("/compressed", handlers.CompressHandler(http.HandlerFunc(getBigResponse)))
	mux.Handle("/hello", handlers.ContentTypeHandler(http.HandlerFunc(getHello), "application/x-www-form-urlencoded"))

	err := http.ListenAndServe(":8181", mux)
	if errors.Is(err, http.ErrServerClosed) { // ErrServerClosed is returned after a call to Shutdown or Close.
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
