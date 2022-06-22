package main

import (
	"log"
	"net/http"
	"os"

	"github.com/krzysztofla/Example.Go.Api/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHelloHandler(logger)

	sm := http.NewServeMux()

	sm.Handle("/hello", hh)

	http.ListenAndServe(":9090", sm)
}
