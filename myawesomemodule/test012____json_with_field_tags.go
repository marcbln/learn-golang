package main

import (
	"encoding/json"
	"fmt"
)

type CustomerWithFieldTags struct {
	ID               int    `json:"id,omitempty"`
	Occupation       string `json:"oc,omitempty"`
	Firstname        string `json:"fn,omitempty"`
	Lastname         string `json:"ln,omitempty"`
	notExportedField string
}

func main() {
	fmt.Printf("example 12\n")

	// create customer struct
	customer := &CustomerWithFieldTags{
		ID:         0,  // omitted
		Occupation: "", // omitted
		// Firstname:        "Dieter", // omitted
		Lastname:         "Müller",
		notExportedField: "aaa",
	}

	// "marshal" (serialize to json)
	marshalled, _ := json.Marshal(customer)
	fmt.Printf("%v %T\n\n", string(marshalled), marshalled) // {"ln":"Müller"} []uint8

	// "unmarshal" (deserialize from json to struct)
	customer2 := &CustomerWithFieldTags{}
	_ = json.Unmarshal(marshalled, customer2)
	fmt.Printf("%v %T\n\n", customer2, customer2) // &{0   Müller } *main.CustomerWithFieldTags

}
