package service

import (
	"time"

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

func (s *WalletService) MonthStatistic(userID int64) (int, float64, error) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstDayOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	firstDayOfMonthStr := firstDayOfMonth.Format("2006-01-02")
	lastDayOfMonthStr := lastDayOfMonth.Format("2006-01-02")

	Wallet, err := s.repos.GetWalletByUserID(userID)
	if err != nil {
		return 0, 0, err
	}

	trn, err := s.repos.MonthStatistic(Wallet.ID, firstDayOfMonthStr, lastDayOfMonthStr)
	if err != nil {
		return 0, 0, err
	}

	var totalAmount float64
	for _, v := range trn {
		totalAmount += v.Amount
	}

	return len(trn), totalAmount, nil
}
