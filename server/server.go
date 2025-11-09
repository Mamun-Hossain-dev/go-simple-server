package server

import (
	"fmt"
	"log"
	"net/http"
	"simple/config"
	"simple/middleware"
	"simple/routes"
	"simple/utils"
)

func StartServer() {
	mux := http.NewServeMux()

	// load config
	cnf := config.LoadConfig()

	// subRouters
	productRouter := routes.ProductRouter()
	helloRouter := routes.HelloRouter()

	// Apply middleware chain per group
	mux.Handle("/api/products/", utils.ChainingMiddlewares(
		http.StripPrefix("/api/products", productRouter),
		middleware.CorsMiddleware,
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
		middleware.AuthMiddleware,
	))

	mux.Handle("/api/hello/", utils.ChainingMiddlewares(
		http.StripPrefix("/api/hello", helloRouter),
		middleware.CorsMiddleware,
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
		middleware.AuthMiddleware,
	))

	fmt.Println("Server is running on port :", cnf.HttpPort)

	log.Fatal(http.ListenAndServe(":"+cnf.HttpPort, mux))
}
