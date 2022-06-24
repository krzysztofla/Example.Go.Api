package data

import (
	"errors"
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

func GetProductById(id int) (*Product, error) {
	for _, item := range productList {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, errors.New("no item with given Id")
}

func AddProduct(p *Product) {
	lp := productList[len(productList)-1]
	p.ID = lp.ID
	p.ID++
	productList = append(productList, p)
}

func UpdateProduct(p *Product) {

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
