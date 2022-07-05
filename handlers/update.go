package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krzysztofla/Example.Go.Api/data"
)

// swagger:route PUT /items products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT Method Invocation - Add New Product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "unable to parse id", http.StatusInternalServerError)
	}
	// fetch the product from the context
	prdt := r.Context().Value(KeyProduct{}).(data.Product)

	data.UpdateProduct(id, &prdt)
}
