package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type MyCustomWrappingError struct {
	Msg       string
	passedArg int
	Code      int
	Err       error
}

// implementation of Error() method of error-interface
func (err *MyCustomWrappingError) Error() string {
	return "msg: " + err.Msg +
		"; code: " + strconv.Itoa(err.Code) +
		"; passedArg: " + strconv.Itoa(err.passedArg) +
		"; original error: " + err.Err.Error()
}

func (e *MyCustomWrappingError) Unwrap() error { return e.Err }

func doSomethingWhichResultsInAnError(giveMe42 int) error {
	if giveMe42 != 42 {
		return &MyCustomWrappingError{
			Msg:       "you passed the wrong argument!",
			passedArg: giveMe42,
			Code:      13,
			Err:       errors.New("some original error"),
		}
	}
	return nil
}

func main() {
	err := doSomethingWhichResultsInAnError(123)
	if err != nil {

		// type assertion
		if e, ok := err.(*MyCustomWrappingError); ok {
			fmt.Println("We had a MyCustomWrappingError here with error code", e.Code)
		}

		log.Fatal(err)
	}
	fmt.Println("will not arrive here")
}
