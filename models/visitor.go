package models

import "time"

type Visitor struct {
	Login          string    `json:"login" binding:"required"`
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
	Patronymic     string    `json:"patronymic" binding:"required"`
	Phone          string    `json:"phone" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	Organization   string    `json:"organization" binding:"required"`
	Note           string    `json:"note" binding:"required"`
	Birthday       time.Time `json:"birthday" binding:"required"`
	PassportSeries string    `json:"passport_series" binding:"required,len=4"`
	PassportNumber string    `json:"passport_number" binding:"required,len=6"`
	Photo          string    `json:"photo" binding:"required"`
}
