package main

import (
	"fmt"
)

func sayHello(c chan string) {
	c <- "Hello Go Channel"
}

func main() {
	c := make(chan string)
	go sayHello(c)
	text := <-c
	fmt.Println(text)
}
