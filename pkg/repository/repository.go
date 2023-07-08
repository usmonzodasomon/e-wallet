package repository

import (
	"github.com/usmonzodasomon/e-wallet/models"
	"gorm.io/gorm"
)

type Authorization interface {
	SignUp(user *models.User) error
	GetUser(phone, password string) (models.User, error)
}

type Wallet interface {
	CreateWallet(wallet *models.Wallet) error
}

type Repository struct {
	Authorization
	Wallet
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Wallet:        NewWalletPostgres(db),
	}
}
