package service

import (
	"github.com/usmonzodasomon/e-wallet/models"
	"github.com/usmonzodasomon/e-wallet/pkg/repository"
)

type WalletService struct {
	repos repository.Wallet
}

func NewWalletService(repos repository.Wallet) *WalletService {
	return &WalletService{repos}
}

func (s *WalletService) CreateWallet(wallet *models.Wallet) error {
	return s.repos.CreateWallet(wallet)
}
