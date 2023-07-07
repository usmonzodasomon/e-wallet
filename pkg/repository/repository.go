package repository

import (
	"github.com/usmonzodasomon/e-wallet/models"
	"gorm.io/gorm"
)

type Authorization interface {
	SignUp(user *models.User) error
	GetUser(phone, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
