package product

import (
	"encoding/json"
	"net/http"
	"simple/utils"
	"strconv"
)

type Handler struct {
	Service ProductService
}

func NewProductHandler(s ProductService) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	data := h.Service.GetAllProducts()
	res := Response{
		Message: "Products fetched successfully!",
		Data:    data,
	}

	utils.SendData(w, res, http.StatusOK)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	createdProduct := h.Service.StoreProduct(p)
	res := ProductResponse{
		Message: "Product created successfully!",
		Data:    createdProduct,
	}

	utils.SendData(w, res, http.StatusCreated)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product Id", http.StatusBadRequest)
		return
	}

	product, err := h.Service.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := ProductResponse{
		Message: "Product fetched successfully!",
		Data:    *product,
	}

	utils.SendData(w, res, http.StatusOK)
}
