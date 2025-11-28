package server

import (
	"fmt"
	"log"
	"net/http"
	"simple/config"
	"simple/internal/product"
	"simple/middleware"
	"simple/routes"

	"simple/utils"
)

func StartServer() {
	mux := http.NewServeMux()

	// load products into the database
	product.LoadFakeProducts()

	// load config
	cfg := config.LoadConfig()

	//------ subRouters-------
	// Apply middleware chain per group
	mux.Handle("/api/products/", utils.ChainingMiddlewares(
		http.StripPrefix("/api/products", routes.ProductRouter()),
		middleware.CorsMiddleware,
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
	))

	mux.Handle("/api/users/", utils.ChainingMiddlewares(
		http.StripPrefix("/api/users", routes.UserRouter(cfg)),
		middleware.CorsMiddleware,
		middleware.LogginMiddleware,
		middleware.TestingMiddleware,
	))

	fmt.Println("Server is running on port :", cfg.HttpPort)

	log.Fatal(http.ListenAndServe(":"+cfg.HttpPort, mux))
}
