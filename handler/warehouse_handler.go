package handler

import (
	"assignment5/models"
	"assignment5/service"
	"assignment5/utils"
	"encoding/json"
	"net/http"
)

type WarehouseHandler struct {
	service service.WarehouseService
}

func NewWarehouseHandler(s service.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{s}
}

func (h *WarehouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.ValidateStruct(warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(warehouse); err != nil {
		http.Error(w, "Failed to create warehouse", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *WarehouseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	warehouses, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get warehouses", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(warehouses)
}
