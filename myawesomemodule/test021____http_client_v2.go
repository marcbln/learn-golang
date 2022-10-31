package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// create a client with some custom configuation
	client := http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 1000 * time.Millisecond,
		},
	}
	res, err := client.Get("https://api.ipify.org")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bites, _ := io.ReadAll(res.Body)
	fmt.Println(string(bites))
}
