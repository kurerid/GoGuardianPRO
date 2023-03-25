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
	IsRegistered(email string) (bool, error)
	GetList() (*models.AccountGetList, error)
}

type Order interface {
	CreatePrivate(input *models.Order) error
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
