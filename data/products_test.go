package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "Kris",
		Price: 1.0,
		SKU:   "asd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
