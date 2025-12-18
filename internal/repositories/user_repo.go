package repositories

import (
	"myapp/internal/models"
	"myapp/pkg/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: database.DB}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) DeleteUser(id []int64) error {
	var user models.User
	err := r.DB.Where("id IN ?", id).Delete(user).Error
	return err
}
