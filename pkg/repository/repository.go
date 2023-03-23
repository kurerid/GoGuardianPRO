package repository

import "github.com/jmoiron/sqlx"

type Auth interface {
}

type Account interface {
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
