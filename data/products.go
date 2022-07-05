package data

import (
	"fmt"
	"time"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU       string `json:"sku" validate:"sku"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

// Products defines a slice of Product
type Products []*Product

// GetProducts returns all products from the database
func GetProducts() []*Product {
	return productList
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductById(id int) (*Product, error) {
	for _, item := range productList {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, ErrProductNotFound
}

func AddProduct(p *Product) {
	lp := productList[len(productList)-1]
	p.ID = lp.ID
	p.ID++
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	prod, err := GetProductById(id)
	if err != nil {
		return ErrProductNotFound
	}

	// todo find id from slice and replace with productList[id] = p
	prod.Name = p.Name
	prod.Price = p.Price
	prod.SKU = p.SKU
	prod.UpdatedAt = time.Now().UTC().String()
	return nil
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
