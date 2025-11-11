package handlers

import (
	"encoding/json"
	"net/http"
	"simple/database"
	"simple/utils"
)

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newuser database.User

	err := json.NewDecoder(r.Body).Decode(&newuser)

	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	createdUser := newuser.Store()
	res := database.UserRes{
		Message: "user created successfully!",
		Data:    createdUser,
	}

	utils.SendData(w, res, http.StatusCreated)
}

// Login user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var logUser database.LoggedUser

	err := json.NewDecoder(r.Body).Decode(&logUser)

	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	user := database.Find(logUser.Email, logUser.Password)
	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	res := database.UserRes{
		Message: "user logged in successfully!",
		Data:    *user,
	}

	utils.SendData(w, res, http.StatusOK)
}
