package routes

import (
	"net/http"
	"simple/handlers"
)

func ProductRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProducts))
	r.Handle("GET /all-products", http.HandlerFunc(handlers.GetProducts))
	r.Handle("GET /{id}", http.HandlerFunc(handlers.GetProductById))

	return r
}
