package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type SaleService interface {
	Create(sale models.Sale) error
	GetAll() ([]models.Sale, error)
}

type saleService struct {
	repo repository.SaleRepository
}

func NewSaleService(r repository.SaleRepository) SaleService {
	return &saleService{r}
}

func (s *saleService) Create(sale models.Sale) error {
	return s.repo.Create(sale)
}

func (s *saleService) GetAll() ([]models.Sale, error) {
	return s.repo.GetAll()
}
