package service

import (
	"assignment5/models"
	"assignment5/repository"
)

type UserService interface {
	Register(user models.User) error
	GetByEmail(email string) (models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Register(user models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByEmail(email string) (models.User, error) {
	return s.repo.GetByEmail(email)
}
