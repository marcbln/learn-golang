package main

import (
	"log"
	"net/http"
)

func createHandler(someParam string) http.Handler {
	f := func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Injected Param: " + someParam))
	}
	return http.HandlerFunc(f)
}

func main() {
	http.Handle("/", createHandler("someParam"))
	log.Fatal(http.ListenAndServe(":8787", nil))
}
