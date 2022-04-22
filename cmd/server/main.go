package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/krzysztofla/Example.Go.Api/handlers"
)

func main() {
	logger := log.New(os.Stdout, "Basket-Api", log.LstdFlags)
	item_handler := handlers.NewItemHandler(logger)

	server_mux := http.NewServeMux()

	server_mux.Handle("/allItems", item_handler)

	http_server := &http.Server{
		Addr:         ":8080",
		Handler:      server_mux,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	http_server.ListenAndServe()
}
