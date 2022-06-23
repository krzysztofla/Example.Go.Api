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

func (h *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	if r.Method == http.MethodGet {
		data, err := json.Marshal(lp)

		if err != nil {
			http.Error(rw, "", http.StatusInternalServerError)
		}
		rw.Write(data)
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}
