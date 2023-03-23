package service

import "GuardianPRO/pkg/repository"

type Auth interface {
}

type Account interface {
}

type Order interface {
}

type Service struct {
	Account
	Order
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:    repos.Auth,
		Account: repos.Account,
		Order:   repos.Order,
	}
}
