package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	inmemoryrepository "github.com/krzysztofla/Example.Go.Api/repository"
)

func GetAllItems(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	items := inmemoryrepository.GetAllItems()
	payload, err := json.Marshal(items)
	if err != nil {
		http.Error(rw, "Unable to pars items payload", http.StatusInternalServerError)
		Logger.Println(err.Error())
	}

	rw.Write(payload)
}

func GetItemById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["id"]

	var sb strings.Builder

	item := inmemoryrepository.GetItemById(itemId)

	if item == nil {
		sb.WriteString("There is no item with coresponding id: ")
		sb.WriteString(itemId)

		http.Error(rw, sb.String(), http.StatusInternalServerError)
		Logger.Println(sb.String())
	}

	res, _ := json.Marshal(item)
	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
}
