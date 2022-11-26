package main

import (
	"log"
	"net/http"
)

func createHandlerV2(someParam string) http.Handler {
	f := func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Injected Param: " + someParam))
	}
	return http.HandlerFunc(f)
}

func createMiddleware(wrapped http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println("before request " + req.URL.Path)
		wrapped.ServeHTTP(rw, req)
		log.Println("after request")
	})
}

func main() {
	http.Handle("/", createMiddleware(createHandlerV2("someParam")))
	log.Fatal(http.ListenAndServe(":8787", nil))
}
