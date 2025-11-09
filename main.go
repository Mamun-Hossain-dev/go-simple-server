package main

import (
	"simple/database"
	"simple/server"
)

func main() {
	server.StartServer()

}

// init function
func init() {
	// Load Products to productlist
	database.LoadProducts()
}
