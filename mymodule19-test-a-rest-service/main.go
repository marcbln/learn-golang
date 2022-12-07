package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

func myHandleFunc(rw http.ResponseWriter, req *http.Request) {
	// rw.WriteHeader(http.StatusOK)
	log.Println(req.URL.Path)
	re := regexp.MustCompile(`^/(.*)\.json$`)
	bIsJsonPath := re.MatchString(req.URL.Path)
	if !bIsJsonPath {
		rw.WriteHeader(http.StatusNotImplemented)
		return
	}
	// ---- build json response
	parts := re.FindStringSubmatch(req.URL.Path)
	name := parts[1]
	xxx := map[string]string{"name": name}
	xxxSerialized, _ := json.Marshal(xxx)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(xxxSerialized)
	rw.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", myHandleFunc)
	http.ListenAndServe(":8383", nil)
}
