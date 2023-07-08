package models

type Wallet struct {
	ID           int64   `json:"id,omitempty" gorm:"primaryKey"`
	UserID       int64   `json:"user_id,omitempty"`
	User         User    `json:"-" gorm:"foreignKey:UserID"`
	Balance      float64 `json:"balance,omitempty" gorm:"default:0"`
	IsIdentified bool    `json:"is_identified,omitempty" gorm:"default:false"`
}
