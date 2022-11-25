package data_test

import (
	"mymodule9-rest-with-validation/data"
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &data.Product{
		Name:        "some Name",
		Description: "some Description",
		Price:       1.23,
		SKU:         "P-XY",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
