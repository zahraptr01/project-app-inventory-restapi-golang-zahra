package handler

import (
	"assignment5/models"
	"assignment5/service"
	"assignment5/utils"
	"encoding/json"
	"net/http"
)

type RackHandler struct {
	service service.RackService
}

func NewRackHandler(s service.RackService) *RackHandler {
	return &RackHandler{s}
}

func (h *RackHandler) Create(w http.ResponseWriter, r *http.Request) {
	var rack models.Rack
	if err := json.NewDecoder(r.Body).Decode(&rack); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.ValidateStruct(rack); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(rack); err != nil {
		http.Error(w, "Failed to create rack", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *RackHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	racks, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get racks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(racks)
}
