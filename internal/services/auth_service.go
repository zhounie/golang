package services

import (
	"myapp/internal/models"
	"myapp/internal/repositories"
)

type AuthService struct {
	Repo *repositories.AuthRepository
}

func NewAuthService() *AuthService {
	return &AuthService{Repo: repositories.NewAuthRepository()}
}

func (s *AuthService) Login(user *models.LoginRequest) (*models.User, error) {

	return s.Repo.FindUserByUserNameAndPassword(user)
}
