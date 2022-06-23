package data

import (
	"time"
)

type Product struct {
	ID          int     `json:"Id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"-"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Products []*Product

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Toy",
		Description: "Funny toy",
		SKU:         "abcd321",
		CreatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Toy",
		Description: "Not funny toy",
		SKU:         "cdas321",
		CreatedAt:   time.Now().UTC().String(),
	},
}
