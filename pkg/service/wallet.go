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

func (s *WalletService) CheckAccount(userID int64) (models.Wallet, error) {
	return s.repos.CheckAccount(userID)
}

func (s *WalletService) GetBalance(userID int64) (float64, error) {
	Wallet, err := s.repos.CheckAccount(userID)
	if err != nil {
		return -1, err
	}
	return s.repos.GetBalance(Wallet.ID)
}

func (s *WalletService) TopUp(TopUp models.TopUp) (models.Transaction, error) {
	return s.repos.TopUp(TopUp)
}
