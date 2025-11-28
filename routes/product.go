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

	// protected routes

	router.HandleFunc("GET /", handler.GetAllProducts)
	router.HandleFunc("GET /{id}", handler.GetProductByID)
	router.Handle("POST /create", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateProduct)))

	return router
}
