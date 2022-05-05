package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	inmemoryrepository "github.com/krzysztofla/Example.Go.Api/repository"
)

func DeleteItem(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["id"]

	var sb strings.Builder

	resp := inmemoryrepository.DeleteItemById(itemId)

	if resp == false {
		sb.WriteString("There is no item with coresponding id: ")
		sb.WriteString(itemId)

		http.Error(rw, sb.String(), http.StatusInternalServerError)
		Logger.Println(sb.String())
	}

	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
}
