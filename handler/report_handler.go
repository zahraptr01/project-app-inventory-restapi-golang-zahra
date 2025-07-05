package handler

import (
	"assignment5/service"
	"encoding/json"
	"net/http"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(s service.ReportService) *ReportHandler {
	return &ReportHandler{s}
}

func (h *ReportHandler) Generate(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("X-Role")

	if role != "admin" {
		http.Error(w, "Forbidden: Only admin can access this report", http.StatusForbidden)
		return
	}

	report, err := h.service.Generate()
	if err != nil {
		http.Error(w, "Failed to generate report", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(report)
}
