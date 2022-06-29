package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/krzysztofla/Example.Go.Api/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProductsHandler(logger)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/items", ph.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.Use(ph.MiddlewareProductValidation)
	postRouter.HandleFunc("/items", ph.AddProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.Use(ph.MiddlewareProductValidation)
	putRouter.HandleFunc("/items/{id}", ph.UpdateProduct)

	server := &http.Server{
		Addr:         ":9091",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Graceful shutdown pattern - it will allow application to finish all the work in the controllers before shutdown. It will log the cause of shutdown.
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel

	logger.Println("Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 40*time.Second)
	server.Shutdown(timeoutContext)
}
