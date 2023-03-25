package service

import (
	"GuardianPRO/models"
	"GuardianPRO/pkg/repository"
)

type Auth interface {
	SignUp(input *models.Account) (*string, error)
	SignIn(input *models.Account) (*string, error)
}

type Account interface {
	IsRegistered(email string) (bool, error)
	GetList() (*models.AccountGetList, error)
}

type Order interface {
	CreatePrivate(input *models.Order) error
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
		Order:   NewOrderService(repos.Order),
	}
}
