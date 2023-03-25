package repository

import (
	"GuardianPRO/models"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) IsRegistered(email string) (bool, error) {
	rows, err := r.db.Query(`SELECT "Email" FROM "Account" WHERE "Email" = $1`, email)
	if err != nil {
		return false, err
	}
	if !rows.Next() {
		return false, nil
	}
	return true, nil
}

func (r *AccountPostgres) GetList() (*models.AccountGetList, error) {
	var output models.AccountGetList
	if err := r.db.Select(&output.List, `SELECT * FROM "Account"`); err != nil {
		return nil, err
	}
	if err := r.db.Get(&output.TotalCount, `WITH templates as
    (SELECT * FROM "Account")
	SELECT count(*) FROM templates`); err != nil {
		return nil, err
	}

	return &output, nil
}
