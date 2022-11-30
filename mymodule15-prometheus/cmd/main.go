package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// the standard http mux will match the / handler to any unregistered path
		// --> check the if the path is really "/"
		if req.URL.Path != "/" { // Check path here
			http.NotFound(rw, req)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("hello"))
	})
	http.HandleFunc("/test-500", func(rw http.ResponseWriter, req *http.Request) {
		// FIXME.. not counted by prometheus exporter.
		rw.WriteHeader(http.StatusInternalServerError)
	})
	http.HandleFunc("/test-503", func(rw http.ResponseWriter, req *http.Request) {
		// FIXME.. not counted by prometheus exporter.
		rw.WriteHeader(http.StatusServiceUnavailable)
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
