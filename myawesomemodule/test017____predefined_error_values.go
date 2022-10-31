package main

import (
	"errors"
	"fmt"
)

var NotFoundError = errors.New("not found")

func someFunc(x int) error {
	if x == 42 {
		return nil
	} else {
		return NotFoundError
	}
}

func main() {
	err := someFunc(1111)
	if err == NotFoundError {
		fmt.Println("The value was not found!")
	}
}
