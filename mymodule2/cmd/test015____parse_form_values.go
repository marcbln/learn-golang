/*
 * parse html form values
 */
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func showForm(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "text/html")
	rw.Write([]byte("<form method=POST>" +
		"<input type=text name=name placeholder=name>" +
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
	rw.Write([]byte("Hello " + name))
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(showForm).Methods("GET")
	r.PathPrefix("/").HandlerFunc(handeForm).Methods("POST")

	http.ListenAndServe(":8282", r)
}
