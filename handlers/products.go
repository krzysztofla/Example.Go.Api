package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krzysztofla/Example.Go.Api/data"
)

type Products struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Method Invocation - Get All Products")
	lp := data.GetProducts()
	data, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
	}
	rw.Write(data)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST Method Invocation - Add New Product")
	encoder := json.NewDecoder(r.Body)
	prdt := &data.Product{}

	encoder.Decode(prdt)

	data.AddProduct(prdt)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT Method Invocation - Add New Product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to parse id", http.StatusInternalServerError)
	}

	encoder := json.NewDecoder(r.Body)
	prdt := &data.Product{}

	encoder.Decode(prdt)

	data.UpdateProduct(id, prdt)
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
