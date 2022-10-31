package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bites, _ := io.ReadAll(res.Body)
	fmt.Println(string(bites))
}
