package handler

import (
	"encoding/json"
	"net/http"

	inmemoryrepository "github.com/krzysztofla/Example.Go.Api/repository"
)

func (i *Item) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	i.l.Debug("Get all records")

	items := inmemoryrepository.GetAllItems()
	payload, error := json.Marshal(items)

	if error != nil {
		http.Error(rw, "Unable to pars items payload", http.StatusInternalServerError)
	}

	rw.Write(payload)
}
