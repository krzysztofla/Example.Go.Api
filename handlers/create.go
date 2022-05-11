package handler

import (
	"net/http"

	inmemoryrepository "github.com/krzysztofla/Example.Go.Api/repository"
)

func CreateItem(rw http.ResponseWriter, r *http.Request) {
	inmemoryrepository.CreateItem(r)

	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
}
