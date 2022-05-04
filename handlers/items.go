package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Keyitem is a key used for the Item object in the context
type KeyItem struct{}

// Items handler for getting and updating items
type Items struct {
	l hclog.Logger
}

// NewItems returns a new items handler with the given logger
func NewItems(l hclog.Logger) *Items {
	return &Items{l}
}

// ErrInvaliditemPath is an error message when the item path is not valid
var ErrInvalidItemsPath = fmt.Errorf("Invalid Path, path should be /item/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getItemID returns the item ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getItemID(r *http.Request) int {
	// parse the item id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
