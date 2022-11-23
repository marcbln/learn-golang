package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
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
