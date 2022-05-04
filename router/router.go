package router

import (
	"github.com/gorilla/mux"
	handler "github.com/krzysztofla/Example.Go.Api/handlers"
)

var RegisterItemStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/items/list", handler.GetAllItems).Methods("GET")
	router.HandleFunc("/items/{id}", handler.GetItemById).Methods("GET")
}
