package middleware

import (
	"context"
	fmt "fmt"
	"log"
	"mymodule9-rest-with-validation/data"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("hello from LoggingMiddleware", r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

type KeyProduct struct{} // this is some kind of constant for the key for the deserialized product in the request context

func ProductDeserializerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// ---- deserialize product from json
		product := &data.Product{}
		err := product.FromJSON(req.Body)
		if err != nil {
			http.Error(rw, "json decoding failed", http.StatusBadRequest)
			return // stop the handler chain
		}

		if err = product.Validate(); err != nil {
			http.Error(rw, fmt.Sprintf("json validation failed: %s", err), http.StatusBadRequest)
			return // stop the handler chain
		}

		// ---- save deserialized product in request context
		ctx := context.WithValue(req.Context(), KeyProduct{}, product)
		req = req.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
