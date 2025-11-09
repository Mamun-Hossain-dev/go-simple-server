package routes

import (
	"net/http"
	"simple/handlers"
)

func HelloRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("GET /", http.HandlerFunc(handlers.HelloHandler))

	return r
}
