package handlers

import (
	"net/http"

	"github.com/krzysztofla/Example.Go.Api/data"
)

// swagger:route POST /items products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST Method Invocation - Add New Product")
	// fetch the product from the context
	prdt := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prdt)
}
