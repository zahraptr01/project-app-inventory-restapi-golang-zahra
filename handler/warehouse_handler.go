package handler

import (
	"assignment5/models"
	"assignment5/service"
	"assignment5/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (h *WarehouseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	warehouse, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(warehouse)
}

func (h *WarehouseHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Update(id, warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *WarehouseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
