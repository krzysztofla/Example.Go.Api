package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krzysztofla/Example.Go.Api/data"
)

type Products struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
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
