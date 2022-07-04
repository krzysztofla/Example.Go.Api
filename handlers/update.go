package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krzysztofla/Example.Go.Api/data"
)

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
