package utils

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, StatusCode int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(StatusCode)
	json.NewEncoder(w).Encode(data)

}
