package main

import (
	"simple/database"
	"simple/server"
)

// global variables
// var ProductList []database.Product

func main() {
	server.StartServer()
}

// init function
func init() {
	// Load Products to productlist
	database.LoadProducts()
}
