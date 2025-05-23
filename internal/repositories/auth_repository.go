package repositories

import (
	"github.com/hnifmaghfur/go-user-service/internal/models"
	r "github.com/hnifmaghfur/go-user-service/internal/requests"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return AuthRepository{db: db}
}

func (ar *AuthRepository) Login(basicAuth r.BasicAuth) (models.User, error) {
	var user models.User
	if err := ar.db.Where("email = ?", basicAuth.Email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ar *AuthRepository) Register(basicAuth r.Register) (models.User, error) {
	user := models.User{
		Email:    basicAuth.Email,
		Password: basicAuth.Password,
		Name:     basicAuth.Name,
	}
	if err := ar.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ar *AuthRepository) UpdatePassword(user models.User, password string) error {
	if err := ar.db.Model(&user).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
