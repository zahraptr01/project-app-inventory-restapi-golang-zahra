package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type WarehouseService interface {
	Create(w models.Warehouse) error
	GetAll() ([]models.Warehouse, error)
	GetByID(id int) (models.Warehouse, error)
	Update(id int, warehouse models.Warehouse) error
	Delete(id int) error
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

func (s *warehouseService) GetByID(id int) (models.Warehouse, error) {
	return s.repo.GetByID(id)
}

func (s *warehouseService) Update(id int, warehouse models.Warehouse) error {
	return s.repo.Update(id, warehouse)
}

func (s *warehouseService) Delete(id int) error {
	return s.repo.Delete(id)
}
