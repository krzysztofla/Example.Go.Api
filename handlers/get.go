package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	inmemoryrepository "github.com/krzysztofla/Example.Go.Api/repository"
)

var logger = log.New(os.Stdout, "Basket-Api", log.LstdFlags)

func GetAllItems(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	items := inmemoryrepository.GetAllItems()
	payload, err := json.Marshal(items)
	if err != nil {
		http.Error(rw, "Unable to pars items payload", http.StatusInternalServerError)
		logger.Println(err.Error())
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
		logger.Println(sb.String())
	}

	res, _ := json.Marshal(item)
	rw.Header().Set("Content-Type", "pkglication/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
}
