package handlers

import (
	"log"
	"net/http"
)

type GoodbyeHandler struct {
	l *log.Logger
}

func NewGoodbyeHandler(l *log.Logger) *GoodbyeHandler {
	return &GoodbyeHandler{l}
}

func (self GoodbyeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	self.l.Println("GoodbyeHandler.ServeHTTP called")
	rw.Write([]byte("Bye!"))
}
