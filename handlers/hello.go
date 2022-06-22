package handlers

import (
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHelloHandler(l *log.Logger) *Hello {
	return &Hello{l: l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Example Handler")
}
