package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krzysztofla/Example.Go.Api/data"
)

// swagger:route GET /items products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Method Invocation - Get All Products")
	lp := data.GetProducts()
	data, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "unable to parse objest", http.StatusInternalServerError)
	}
	rw.Write(data)
}

// swagger:route GET /items/{id} product getSingle
// Return prodcuct with given id from the database
// responses:
//	200: productResponse

// ListAll handles GET requests and returns all current products
func (p *Products) GetProductById(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET By Id Method Invocation - Get All Products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "unable to parse id", http.StatusInternalServerError)
	}

	prdt, err := data.GetProductById(id)

	if err != nil {
		http.Error(rw, "no item with given id", http.StatusInternalServerError)
	}

	item, err := json.Marshal(prdt)
	if err != nil {
		http.Error(rw, "unable to parse objest", http.StatusInternalServerError)
	}
	rw.Write(item)
}
