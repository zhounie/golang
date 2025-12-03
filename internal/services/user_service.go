package services

import (
	"myapp/internal/models"
	"myapp/internal/repositories"
	"errors"
	"log"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{Repo: repositories.NewUserRepository()}
}

func (s *UserService) CreateUser(user *models.User) error {
	log.Println(user)
	if len(user.Password) < 5 {
		return errors.New("password is len min 5")
	}
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.FindUserByID(id)
}