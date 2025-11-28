package user

import (
	"encoding/json"
	"net/http"
	"simple/config"
	"simple/utils"
)

type Handler struct {
	service Service
	cfg     *config.Config
}

func NewUserHandler(s Service, cfg *config.Config) *Handler {
	return &Handler{
		service: s,
		cfg:     cfg,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	createdUser := h.service.RegisterUser(u)

	res := UserRes{
		Message: "User created successfully!",
		Data:    createdUser,
	}

	utils.SendData(w, res, http.StatusCreated)
}

func (h *Handler) Loggin(w http.ResponseWriter, r *http.Request) {
	var logUser LoggedUser

	if err := json.NewDecoder(r.Body).Decode(&logUser); err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	user, err := h.service.LoginUser(logUser.Email, logUser.Password)
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

	// jwt create
	accessToken, _ := utils.CreateAccessToken(user.ID, []byte(h.cfg.AccessSecretkey))
	refreshToken, _ := utils.CreateRefreshToken(user.ID, []byte(h.cfg.RefreshSecretKey))

	// ----> secure cookie set
	utils.SetCookie(w, "access_token", accessToken, 15*60)
	utils.SetCookie(w, "refresh_token", refreshToken, 30*24*60*60)

	res := LoginResponse{
		Message:      "User logged in successfully!",
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.SendData(w, res, http.StatusOK)
}
