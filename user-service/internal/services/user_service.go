package services

import (
	"micro-service-go/internal/models"
	"micro-service-go/internal/repositories"
)

type UserService interface {
	GetUsers() ([]models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) GetUsers() ([]models.User, error) {
	return u.userRepo.GetUsers()
}
