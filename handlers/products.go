package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"simple/database"
	"simple/utils"
	"strconv"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var product database.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	product.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product created successfully!",
		"data":    product,
	})
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.PathValue("id")); err != nil {
		http.Error(w, "Invalid product Id", http.StatusBadRequest)
		return
	} else {
		log.Println("Requested product ID", id)

		for _, p := range database.ProductList {
			if p.ID == id {
				utils.SendData(w, database.ProductResponse{
					Message: "Product fetched successfully!",
					Data:    p,
				}, http.StatusOK)
				return
			}

		}
		http.Error(w, "Product not found", http.StatusNotFound)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(database.Response{
		Message: "Products fetched successfully!",
		Data:    database.ProductList,
	})
}
