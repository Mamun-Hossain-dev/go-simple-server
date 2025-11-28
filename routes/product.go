package routes

import (
	"net/http"
	"simple/internal/product"
	"simple/middleware"
)

func ProductRouter() *http.ServeMux {
	router := http.NewServeMux()

	repo := product.NewProductRepository()
	service := product.NewProductService(repo)
	handler := product.NewProductHandler(service)

	// PUBLIC ROUTES
	router.HandleFunc("GET /", handler.GetProducts)
	router.HandleFunc("GET /{id}", handler.GetProduct)

	// PROTECTED ROUTES (POST, PUT, DELETE)
	router.Handle("POST /",
		middleware.AuthMiddleware(http.HandlerFunc(handler.CreateProduct)),
	)

	router.Handle("PUT /{id}",
		middleware.AuthMiddleware(http.HandlerFunc(handler.UpdateProduct)),
	)

	router.Handle("DELETE /{id}",
		middleware.AuthMiddleware(http.HandlerFunc(handler.DeleteProduct)),
	)

	return router
}
