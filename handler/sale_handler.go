package handler

import (
	"assignment5/models"
	"assignment5/service"
	"assignment5/utils"
	"encoding/json"
	"net/http"
)

type SaleHandler struct {
	service service.SaleService
}

func NewSaleHandler(s service.SaleService) *SaleHandler {
	return &SaleHandler{s}
}

func (h *SaleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var sale models.Sale
	if err := json.NewDecoder(r.Body).Decode(&sale); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.ValidateStruct(sale); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(sale); err != nil {
		http.Error(w, "Failed to create sale", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *SaleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	sales, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get sales", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sales)
}
