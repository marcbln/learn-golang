package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
)

const PORT = 8181

type Customer13 struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var customer *Customer13

/*
 * http GET http://localhost:8181/customer
 */
func myGetFunc(rw http.ResponseWriter, req *http.Request) {
	bites, _ := json.Marshal(customer)
	rw.Header().Add("content-type", "application/json")
	rw.Write(bites)
}

/*
 * echo '{"firstname": "Marc", "lastname": "Christenfeldt", "id": 123}' | http POST http://localhost:8181/customer
 */
func myPostFunc(rw http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	json.Unmarshal(body, customer)
	rw.WriteHeader(http.StatusNoContent)
}

func main() {
	fmt.Printf("example 13\n")

	// create customer struct
	customer = &Customer13{
		ID:        1,
		Firstname: "Hans",
		Lastname:  "Dieter",
	}

	r := mux.NewRouter()
	r.HandleFunc("/customer", myGetFunc).Methods("GET")
	r.HandleFunc("/customer", myPostFunc).Methods("POST")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), r)
	if err != nil {
		fmt.Printf("error starting web server on port %v: %v", PORT, err)
		os.Exit(1)
	}
}
