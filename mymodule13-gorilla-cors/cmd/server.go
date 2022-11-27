package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", func(rw http.ResponseWriter, req *http.Request) {
		log.Println("path: " + req.URL.Path)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("hello from " + req.RequestURI))
	})

	ch := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:8282"}))

	log.Println("starting backend server on :8181")
	log.Fatal(http.ListenAndServe(":8181", ch(r)))
}
