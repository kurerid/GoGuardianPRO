package service

import (
	"GuardianPRO/models"
	"GuardianPRO/pkg/repository"
)

type Auth interface {
	SignUp(account *models.Account) (*string, error)
	SignIn(account *models.Account) (*string, error)
}

type Account interface {
	IsRegistered(login string) (bool, error)
	GetList() (*models.AccountGetList, error)
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
		Auth:    NewAuthService(repos.Auth),
		Account: NewAccountService(repos.Account),
		Order:   repos.Order,
	}
}
