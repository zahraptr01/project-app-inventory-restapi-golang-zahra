package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type ItemService interface {
	Create(item models.Item) error
	GetAll() ([]models.Item, error)
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) ItemService {
	return &itemService{r}
}

func (s *itemService) Create(item models.Item) error {
	return s.repo.Create(item)
}

func (s *itemService) GetAll() ([]models.Item, error) {
	return s.repo.GetAll()
}
