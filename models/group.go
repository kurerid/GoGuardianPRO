package models

type Group struct {
	ID       int       `json:"id" binding:"required"`
	Order    int       `json:"order" binding:"required"`
	Accounts []Account `json:"accounts" binding:"required"`
}
