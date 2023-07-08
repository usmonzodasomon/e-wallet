package service

import (
	"github.com/usmonzodasomon/e-wallet/models"
	"github.com/usmonzodasomon/e-wallet/pkg/repository"
)

type Authorization interface {
	SignUp(user *models.User) error
	GenerateToken(user models.User) (string, error)
	ParseToken(accessToken string) (int64, error)
}

type Wallet interface {
	CreateWallet(wallet *models.Wallet) error
	CheckAccount(userID int64) (models.Wallet, error)
	GetBalance(userID int64) (float64, error)
}

type Service struct {
	Authorization
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Wallet:        NewWalletService(repos.Wallet),
	}
}
