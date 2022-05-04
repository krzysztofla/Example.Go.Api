package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krzysztofla/Example.Go.Api/router"
)

func main() {
	mux_router := mux.NewRouter()

	router.RegisterItemStoreRoutes(mux_router)

	http.ListenAndServe("localhost:8080", mux_router)
}
