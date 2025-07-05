package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type CategoryService interface {
	Create(category models.Category) error
	GetAll() ([]models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{r}
}

func (s *categoryService) Create(category models.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}
