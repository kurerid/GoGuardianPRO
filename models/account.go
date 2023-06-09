package models

type Account struct {
	Email    string `json:"email" db:"Email" binding:"required"`
	Password string `json:"password" db:"Password" binding:"required,min=8"`
}

type AccountGetList struct {
	List       []Account `json:"list"`
	TotalCount uint      `json:"total_count"`
}
