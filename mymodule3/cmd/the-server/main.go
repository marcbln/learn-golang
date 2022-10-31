package main

import (
	"github.com/gorilla/mux"
	registrierung "mymodule3"
	"net/http"
)

func main() {
	regHandler := &registrierung.RegistrierungsHandler{}
	r := mux.NewRouter()
	r.PathPrefix("/").Methods("POST").Handler(regHandler)
	r.PathPrefix("/").Methods("GET").HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		rw.Write([]byte("Try POST method..."))
	})
	http.ListenAndServe(":8383", r)
}
