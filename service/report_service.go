package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type ReportService interface {
	Generate() (models.Report, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(r repository.ReportRepository) ReportService {
	return &reportService{r}
}

func (s *reportService) Generate() (models.Report, error) {
	return s.repo.GenerateReport()
}
