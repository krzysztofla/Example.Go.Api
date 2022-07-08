// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/krzysztofla/Example.Go.Api/data"

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All current products
	// in: body
	Body []data.Product
}

// single product
// swagger: response productResponse
type producstResponseWrapper struct {
	// Producst with corresponding id
	// in: body
	Body data.Product
}
