package service

import "github.com/usmonzodasomon/e-wallet/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
