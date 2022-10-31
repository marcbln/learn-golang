package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	ID               int
	Firstname        string
	Lastname         string
	notExportedField string
}

func main() {
	fmt.Printf("example 11\n")

	// create customer struct
	customer := &Customer{
		ID:               1,
		Firstname:        "Hans äöü",
		Lastname:         "Dieter",
		notExportedField: "aaa",
	}

	// "marshal" (serialize to json)
	marshalled, _ := json.Marshal(customer)
	fmt.Printf("%v %T\n\n", string(marshalled), marshalled) // {"ID":1,"Firstname":"Hans äöü","Lastname":"Dieter"} []uint8

	// "unmarshal" (deserialize from json to struct)
	customer2 := &Customer{}
	_ = json.Unmarshal(marshalled, customer2)
	fmt.Printf("%v %T\n\n", customer2, customer2) // &{1 Hans äöü Dieter} *main.Customer

}
