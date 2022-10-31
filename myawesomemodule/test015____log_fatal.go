package main

import (
	"errors"
	"fmt"
	"log"
)

func createSomeError() error {
	return errors.New("some error")
}
func createSomeOtherError() error {
	return fmt.Errorf("some other error %v %v", 1, 2)
}

func main() {
	if err := createSomeError(); err != nil {
		fmt.Println("Fehler:", err)
	}

	defer fmt.Println("some deferred code")

	if err := createSomeOtherError(); err != nil {
		// The log message goes to the configured log output.
		// calls os.Exit - exits immediately - deferred functions can't be run.
		// in general: Fatal should be avoided in favor of Panic. Only use Fatal if you know for sure there's no defer in the code
		log.Fatalln("some fatal error:", err)
	}

	fmt.Println("should never reach here .. as log.Fatalln exits code execution")
}
