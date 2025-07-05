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

func (h *RackHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	rack, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(rack)
}

func (h *RackHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var rack models.Rack
	if err := json.NewDecoder(r.Body).Decode(&rack); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Update(id, rack); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *RackHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
