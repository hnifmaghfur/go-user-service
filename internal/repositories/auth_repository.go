package repositories

import (
	"github.com/hnifmaghfur/go-user-service/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return AuthRepository{db: db}
}

func (ar *AuthRepository) Login(email string) (models.User, error) {
	var user models.User
	if err := ar.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
