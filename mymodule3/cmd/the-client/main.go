package main

/*
 * the client posting registration form data to our registration service
 */

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{
		"Firstname":     {"Hans"},
		"Lastname":      {"Dieter"},
		"TermsAccepted": {"true"},
	}
	res, err := http.PostForm("http://localhost:8383", data)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode, res.Status)
}
