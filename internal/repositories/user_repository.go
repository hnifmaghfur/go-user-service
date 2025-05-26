package repositories

import (
	"github.com/hnifmaghfur/go-user-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (ur *UserRepository) Post(user models.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
