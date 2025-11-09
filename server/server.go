package server

import (
	"fmt"
	"log"
	"net/http"
	"simple/handlers"
	"simple/middleware"
	"simple/routes"
	"simple/utils"
)

func StartServer() {
	mux := http.NewServeMux()

	// subRouters
	productRouter := routes.ProductRouter()

	// Apply middleware chain per group
	mux.Handle("/api/products/", utils.ChainingMiddlewares(
		productRouter,
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
	))

	mux.Handle("/hello", utils.ChainingMiddlewares(http.HandlerFunc(handlers.HelloHandler),
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
		middleware.AuthMiddleware,
	))

	// mux.Handle("/hello", middleware.LogginMiddleware(http.HandlerFunc(routes.HelloHandler)))
	// mux.Handle("GET /products", middleware.TestingMiddleware(middleware.LogginMiddleware(http.HandlerFunc(routes.GetProducts))))
	// mux.Handle("POST /create-product", middleware.TestingMiddleware(middleware.LogginMiddleware(http.HandlerFunc(routes.CreateProducts))))
	// mux.Handle("GET /products/{id}", middleware.TestingMiddleware(middleware.LogginMiddleware(http.HandlerFunc(routes.GetProductById))))

	mux.Handle("GET /products", utils.ChainingMiddlewares(
		http.HandlerFunc(handlers.GetProducts),
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
		middleware.AuthMiddleware,
	))

	mux.Handle("GET /products/{id}", utils.ChainingMiddlewares(
		http.HandlerFunc(handlers.GetProductById),
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
		middleware.AuthMiddleware,
	))

	mux.Handle("POST /products", utils.ChainingMiddlewares(
		http.HandlerFunc(handlers.CreateProducts),
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
	))

	// wrap entire mux with middleware (global)
	handler := middleware.CorsMiddleware(mux)

	fmt.Println("Server is running on port :8080")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
