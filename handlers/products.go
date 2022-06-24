package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		rgx := regexp.MustCompile(`/([0-9]+)`)
		path := r.URL.Path
		idg := rgx.FindAllStringSubmatch(path, -1)

		if len(idg) != 1 {
			http.Error(rw, "Invalid request", http.StatusBadRequest)
			return
		}

		if len(idg[0]) != 2 {
			http.Error(rw, "Invalid request", http.StatusBadRequest)
			return
		}

		idString := idg[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, " Invalid Request", http.StatusBadRequest)
			return
		}
		p.l.Println(id)
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

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

}
