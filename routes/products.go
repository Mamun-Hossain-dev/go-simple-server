package routes

import (
	"net/http"
	"simple/handlers"
)

func ProductRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("GET /all-products", http.HandlerFunc(handlers.GetProducts))
	r.Handle("POST /create-product", http.HandlerFunc(handlers.CreateProducts))

	r.Handle("GET /{id}", http.HandlerFunc(handlers.GetProductById)) // regex only numbers

	return r
}
