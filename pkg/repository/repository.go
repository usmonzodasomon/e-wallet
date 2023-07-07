package repository

type Authorization interface {
}

type Wallet interface {
}

type Repository struct {
	Authorization
	Wallet
}

func NewRepository() *Repository {
	return &Repository{}
}
