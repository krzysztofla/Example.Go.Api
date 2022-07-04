package handlers

import (
	"net/http"

	"github.com/krzysztofla/Example.Go.Api/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST Method Invocation - Add New Product")
	prdt := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prdt)
}
