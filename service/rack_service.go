package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type RackService interface {
	Create(r models.Rack) error
	GetAll() ([]models.Rack, error)
}

type rackService struct {
	repo repository.RackRepository
}

func NewRackService(r repository.RackRepository) RackService {
	return &rackService{r}
}

func (s *rackService) Create(r models.Rack) error {
	return s.repo.Create(r)
}

func (s *rackService) GetAll() ([]models.Rack, error) {
	return s.repo.GetAll()
}
