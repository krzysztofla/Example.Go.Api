package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krzysztofla/Example.Go.Api/models"
)

type ItemHandler struct {
	l *log.Logger
}

func NewItemHandler(l *log.Logger) *ItemHandler {
	return &ItemHandler{l}
}

func (g *ItemHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	items := models.GetAllItems()
	payload, error := json.Marshal(items)

	if error != nil {
		http.Error(rw, "Unable to pars items payload", http.StatusInternalServerError)
	}

	rw.Write(payload)
}
