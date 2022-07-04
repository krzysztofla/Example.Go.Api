package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/krzysztofla/Example.Go.Api/data"
)

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Method Invocation - Get All Products")
	lp := data.GetProducts()
	data, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
	}
	rw.Write(data)
}
