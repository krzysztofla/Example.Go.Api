package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "Basket-Api", log.LstdFlags)
	get_handler := get.GetAllHandler(logger)

	server_mux := http.NewServeMux()

	server_mux.Handle("/allItems", get_handler)
	server_mux.Handle("/items", get_handler)

	http_server := &http.Server{
		Addr:         ":8080",
		Handler:      server_mux,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	http_server.ListenAndServe()
}
