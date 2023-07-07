package repository

import (
	"github.com/usmonzodasomon/e-wallet/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) SignUp(user *models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUser(phone string, password string) (models.User, error) {
	var user models.User
	if err := r.db.Where("phone = ? AND password = ?", phone, password).Take(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
