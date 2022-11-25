/*
 * parse html form values
 */
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	"log"
	"net/http"
	"strings"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func showForm(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "text/html")
	rw.Write([]byte("<form method=POST>" +
		"<div><input type=text name=name placeholder=name></div>" +
		"<div><input type=email name=email placeholder=email></div>" +
		"<button type=submit>go!</button>" +
		"</form>"))
}

func handeForm(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	name := req.Form.Get("name")
	rw.Write([]byte("Hello " + name + "\n\n"))

	// ---- output all form inputs
	for key, val := range req.Form {
		rw.Write([]byte(key + "=" + strings.Join(val, ", ") + "\n"))
	}
}

func main() {
	err := env.Parse()
	if err != nil {
		log.Fatalln("error parsing env: ", err)
	}

	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(showForm).Methods("GET")
	r.PathPrefix("/").HandlerFunc(handeForm).Methods("POST")

	fmt.Println("listening on ", *bindAddress)
	if err := http.ListenAndServe(*bindAddress, r); err != nil {
		log.Fatal(err)
	}
}
