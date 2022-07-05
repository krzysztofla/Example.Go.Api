package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "Kris",
		Price: 1.0,
		SKU:   "asd",
	}
	validation := NewValidation()
	err := validation.Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}
