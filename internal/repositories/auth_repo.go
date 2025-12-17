package repositories

import (
	"errors"
	"myapp/internal/models"
	"myapp/pkg/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{DB: database.DB}
}

func (r *AuthRepository) FindUserByUserNameAndPassword(loginReq *models.LoginRequest) (*models.User, error) {
	var user models.User
	result := r.DB.Where("username = ?", loginReq.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码不正确")
		}
		return nil, result.Error
	}
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(loginReq.Password),
	)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, errors.New("用户名或密码不正确")
		}
		return nil, errors.New("密码校验失败")
	}
	return &user, nil
}
