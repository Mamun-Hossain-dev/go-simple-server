package routes

import (
	"net/http"
	"simple/config"
	"simple/internal/user"
)

func UserRouter(cfg *config.Config) *http.ServeMux {
	router := http.NewServeMux()

	repo := user.NewUserRepository()
	service := user.NewUserService(repo)
	handler := user.NewUserHandler(service, cfg)

	router.HandleFunc("POST /register", handler.CreateUser)
	router.HandleFunc("POST /login", handler.Loggin)

	return router
}
