package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type CategoryService interface {
	Create(category models.Category) error
	GetAll() ([]models.Category, error)
	GetByID(id int) (models.Category, error)
	Update(id int, category models.Category) error
	Delete(id int) error
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

func (s *categoryService) GetByID(id int) (models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryService) Update(id int, category models.Category) error {
	return s.repo.Update(id, category)
}

func (s *categoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
