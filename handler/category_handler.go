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

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{s}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := utils.ValidateStruct(cat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(cat); err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Update(id, category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
