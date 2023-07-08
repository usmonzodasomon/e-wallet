package models

type Transaction struct {
	ID              int64   `json:"id,omitempty" gorm:"primaryKey"`
	SenderID        int64   `json:"sender_id,omitempty"`
	ReceiverID      int64   `json:"recipient_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	Status          string  `json:"status,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
}

type TopUp struct {
	SenderID      int64   `json:"sender_id,omitempty"`
	ReceiverPhone string  `json:"receiver_phone,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
}
