package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HelloHandler struct {
	l *log.Logger
}

func NewHelloHandler(l *log.Logger) *HelloHandler {
	return &HelloHandler{l: l}
}

func (self *HelloHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	self.l.Println("HelloHandler.ServeHTTP called")
	d, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s!", d)

}
