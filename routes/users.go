package routes

import (
	"net/http"
	"simple/handlers"
)

func UserRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("POST /", http.HandlerFunc(handlers.CreateUser))
	r.Handle("POST /login", http.HandlerFunc(handlers.LoginUser))

	return r
}
