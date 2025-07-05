package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type UserService interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
	GetByID(id int) (models.User, error)
	Update(id int, user models.User) error
	Delete(id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(user models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetByID(id int) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Update(id int, user models.User) error {
	return s.repo.Update(id, user)
}

func (s *userService) Delete(id int) error {
	return s.repo.Delete(id)
}
