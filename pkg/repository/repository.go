package repository

import (
	"GuardianPRO/models"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Auth
	Account
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:    NewAuthPostgres(db),
		Account: NewAccountPostgres(db),
		Order:   NewOrderPostgres(db),
	}
}
