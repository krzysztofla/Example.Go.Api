package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/krzysztofla/Example.Go.Api/data"
)

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prdt := data.Product{}
		encoder := json.NewDecoder(r.Body)

		encoder.Decode(&prdt)
		validation := data.NewValidation()

		err := validation.Validate(prdt)
		if err != nil {
			http.Error(rw, "Validation error. Please make sure all properties are ok", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prdt)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
