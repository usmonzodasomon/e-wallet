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

func (r *WalletPostgres) CheckAccount(userID int64) (models.Wallet, error) {
	var Wallet models.Wallet
	if err := r.db.Where("user_id = ?", userID).Take(&Wallet).Error; err != nil {
		return models.Wallet{}, err
	}
	return Wallet, nil
}

func (r *WalletPostgres) GetBalance(walletID int64) (float64, error) {
	var balance float64
	if err := r.db.Model(models.Wallet{}).Select("balance").Where("id = ?", walletID).Scan(&balance).Error; err != nil {
		return -1, err
	}
	return balance, nil
}
