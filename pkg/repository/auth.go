package repository

import (
	"GuardianPRO/models"
	"errors"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignUp(account *models.Account) (*string, error) {
	_, err := r.db.Exec(`INSERT INTO "Account"("Email", "Password") VALUES($1,$2)`, account.Email, account.Password)
	if err != nil {
		return nil, err
	}
	auth := "info for authorization header"
	return &auth, nil
}

func (r *AuthPostgres) SignIn(account *models.Account) (*string, error) {
	var expectedAccount models.Account
	err := r.db.Get(&expectedAccount, `SELECT * FROM "Account" WHERE "Email" = $1 AND "Password" = $2`, account.Email, account.Password)
	if err != nil {
		return nil, err
	}
	if expectedAccount.Password != account.Password {
		return nil, errors.New("Invalid password")
	}
	auth := "info for authorization header"
	return &auth, nil
}
