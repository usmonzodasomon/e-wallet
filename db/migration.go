package db

import (
	"github.com/usmonzodasomon/e-wallet/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, models.Wallet{}, models.Transaction{}); err != nil {
		return err
	}
	return nil
}
