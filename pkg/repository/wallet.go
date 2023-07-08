package repository

import (
	"github.com/usmonzodasomon/e-wallet/models"
	"gorm.io/gorm"
)

type WalletPostgres struct {
	db *gorm.DB
}

func NewWalletPostgres(db *gorm.DB) *WalletPostgres {
	return &WalletPostgres{db}
}

func (r *WalletPostgres) CreateWallet(wallet *models.Wallet) error {
	if err := r.db.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}
