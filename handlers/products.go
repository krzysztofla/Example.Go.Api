package handlers

import (
	"context"
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
	prdt := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prdt)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT Method Invocation - Add New Product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "unable to parse id", http.StatusInternalServerError)
	}
	prdt := r.Context().Value(KeyProduct{}).(data.Product)

	data.UpdateProduct(id, &prdt)
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prdt := data.Product{}
		encoder := json.NewDecoder(r.Body)

		encoder.Decode(&prdt)

		ctx := context.WithValue(r.Context(), KeyProduct{}, prdt)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
