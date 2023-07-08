package repository

import (
	"errors"

	"github.com/usmonzodasomon/e-wallet/models"
	"gorm.io/gorm"
)

type WalletPostgres struct {
	db *gorm.DB
}

func NewWalletPostgres(db *gorm.DB) *WalletPostgres {
	return &WalletPostgres{db}
}

const (
	successfully   = "успешно"
	unsuccessfully = "неуспешно"
	replenishment  = "пополнение"
	subtraction    = "вычитывание"
)

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

func (r *WalletPostgres) GetWalletByPhoneNumber(phoneNumber string) (models.Wallet, error) {
	var User models.User
	if err := r.db.Where("phone = ?", phoneNumber).Take(&User).Error; err != nil {
		return models.Wallet{}, err
	}

	var Wallet models.Wallet
	if err := r.db.Where("user_id = ?", User.ID).Take(&Wallet).Error; err != nil {
		return models.Wallet{}, err
	}
	return Wallet, nil
}

func (r *WalletPostgres) GetWalletByUserID(userID int64) (models.Wallet, error) {
	var User models.User
	if err := r.db.Where("id = ?", userID).Take(&User).Error; err != nil {
		return models.Wallet{}, err
	}

	var Wallet models.Wallet
	if err := r.db.Where("user_id = ?", User.ID).Take(&Wallet).Error; err != nil {
		return models.Wallet{}, err
	}
	return Wallet, nil
}

func (r *WalletPostgres) TopUp(TopUp models.TopUp) (models.Transaction, error) {
	receiverWallet, err := r.GetWalletByPhoneNumber(TopUp.ReceiverPhone)
	if err != nil {
		return models.Transaction{}, err
	}

	senderWallet, err := r.GetWalletByUserID(TopUp.SenderID)
	if err != nil {
		return models.Transaction{}, err
	}

	if senderWallet.Balance < TopUp.Amount {
		return models.Transaction{}, errors.New("недостаточно денег на балансе")
	}

	if !receiverWallet.IsIdentified && receiverWallet.Balance+TopUp.Amount > 10000 {
		return models.Transaction{}, errors.New("для неидентифицированного кошелька баланс не может превышать 10000")
	} else if receiverWallet.IsIdentified && receiverWallet.Balance+TopUp.Amount > 100000 {
		return models.Transaction{}, errors.New("баланс не может превышать 100000")
	}

	tx := r.db.Begin()
	senderWallet.Balance -= TopUp.Amount
	if err := tx.Save(&senderWallet).Error; err != nil {
		r.CreateTransaction(tx, models.Transaction{
			SenderID:        senderWallet.ID,
			ReceiverID:      receiverWallet.ID,
			Amount:          TopUp.Amount,
			Status:          unsuccessfully,
			TransactionType: replenishment,
		})
		tx.Rollback()
		return models.Transaction{}, err
	}
	receiverWallet.Balance += TopUp.Amount
	if err := tx.Save(&receiverWallet).Error; err != nil {
		r.CreateTransaction(tx, models.Transaction{
			SenderID:        senderWallet.ID,
			ReceiverID:      receiverWallet.ID,
			Amount:          TopUp.Amount,
			Status:          unsuccessfully,
			TransactionType: subtraction,
		})
		tx.Rollback()
		return models.Transaction{}, err
	}

	transactionSender := models.Transaction{
		SenderID:        senderWallet.ID,
		ReceiverID:      receiverWallet.ID,
		Amount:          TopUp.Amount,
		Status:          successfully,
		TransactionType: replenishment,
	}
	if err := r.CreateTransaction(tx, transactionSender); err != nil {
		return models.Transaction{}, err
	}

	transactionReceiver := models.Transaction{
		SenderID:        senderWallet.ID,
		ReceiverID:      receiverWallet.ID,
		Amount:          TopUp.Amount,
		Status:          successfully,
		TransactionType: subtraction,
	}
	if err := r.CreateTransaction(tx, transactionReceiver); err != nil {
		return models.Transaction{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.Transaction{}, err
	}

	return transactionSender, nil
}

func (r *WalletPostgres) CreateTransaction(db *gorm.DB, transaction models.Transaction) error {
	return db.Create(&transaction).Error
}

func (r *WalletPostgres) MonthStatistic(walletID int64, firstDayMonth string, lastDayMonth string) (trn []models.Transaction, err error) {
	if err := r.db.Where("receiver_id = ? AND date >= ? AND date <= ? AND transaction_type = ?", walletID, firstDayMonth, lastDayMonth, replenishment).Find(&trn).Error; err != nil {
		return nil, err
	}
	return trn, nil
}
