package main

import (
	"errors"
	"fmt"
)

var NotFoundError2 = errors.New("not found")

func someFunc2(x int) error {
	if x == 42 {
		return nil
	} else {
		err := NotFoundError2
		return fmt.Errorf("FAIL: %w", err) // creates an error implementing the Unwrap function
	}
}

func main() {
	err := someFunc2(1111)
	if err == NotFoundError2 {
		// not executed:
		fmt.Println("The value was not found ... V1")
	}
	if errors.Is(err, NotFoundError2) {
		// executed:
		fmt.Println("The value was not found ... V2")
	}
}
