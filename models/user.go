package models

type User struct {
	ID       int64  `json:"id,omitempty" gorm:"primaryKey"`
	FIO      string `json:"fio,omitempty"`
	Age      int    `json:"age,omitempty"`
	Phone    string `json:"phone,omitempty" gorm:"not null, unique"`
	Password string `json:"password,omitempty" gorm:"not null"`
}
