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
	CheckAccount(userID int64) (models.Wallet, error)
	GetBalance(walletID int64) (float64, error)
	GetWalletByPhoneNumber(phone string) (models.Wallet, error)
	GetWalletByUserID(userID int64) (models.Wallet, error)
	TopUp(TopUp models.TopUp) (models.Transaction, error)
	CreateTransaction(db *gorm.DB, transaction models.Transaction) error
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
