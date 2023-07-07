package models

type User struct {
	ID       int64  `json:"id,omitempty"`
	FIO      string `json:"fio,omitempty"`
	Age      int    `json:"age,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}
