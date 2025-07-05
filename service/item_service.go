package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type ItemService interface {
	Create(item models.Item) error
	GetAll() ([]models.Item, error)
	GetByID(id int) (models.Item, error)
	Update(id int, item models.Item) error
	Delete(id int) error
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

func (s *itemService) GetByID(id int) (models.Item, error) {
	return s.repo.GetByID(id)
}

func (s *itemService) Update(id int, item models.Item) error {
	return s.repo.Update(id, item)
}

func (s *itemService) Delete(id int) error {
	return s.repo.Delete(id)
}
