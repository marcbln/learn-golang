package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func myHandlerFunc3(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.Write([]byte(fmt.Sprintf("... %v %T", vars, vars)))
}

func main() {
	fmt.Printf("registering routes...\n")
	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(myHandlerFunc3)
	router.HandleFunc("/products/{id}", myHandlerFunc3) // with path variable
	router.Methods("GET")                               // order is important! calling the before the route registration, it does not work (404)

	fmt.Printf("starting server...\n")
	http.ListenAndServe(":8080", router)
}
