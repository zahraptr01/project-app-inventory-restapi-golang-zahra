package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type RackService interface {
	Create(r models.Rack) error
	GetAll() ([]models.Rack, error)
	GetByID(id int) (models.Rack, error)
	Update(id int, rack models.Rack) error
	Delete(id int) error
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

func (s *rackService) GetByID(id int) (models.Rack, error) {
	return s.repo.GetByID(id)
}

func (s *rackService) Update(id int, rack models.Rack) error {
	return s.repo.Update(id, rack)
}

func (s *rackService) Delete(id int) error {
	return s.repo.Delete(id)
}
