package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func myHandlerFunc2(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from http server äöüßα"))
}

func main() {
	fmt.Printf("registering routes...\n")
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(myHandlerFunc2)
	router.Methods("GET") // order is important! calling the before the route registration, it does not work (404)

	fmt.Printf("starting server...\n")
	http.ListenAndServe(":8080", router)
}
