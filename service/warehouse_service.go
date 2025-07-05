package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type WarehouseService interface {
	Create(w models.Warehouse) error
	GetAll() ([]models.Warehouse, error)
}

type warehouseService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseService(r repository.WarehouseRepository) WarehouseService {
	return &warehouseService{r}
}

func (s *warehouseService) Create(w models.Warehouse) error {
	return s.repo.Create(w)
}

func (s *warehouseService) GetAll() ([]models.Warehouse, error) {
	return s.repo.GetAll()
}
