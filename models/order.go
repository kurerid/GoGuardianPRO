package models

import "time"

const (
	PRIVATE_ORDER = "private"
	GROUP_ORDER   = "group"
)

type Order struct {
	ID           int       `json:"id" binding:"required"`
	BeginDate    time.Time `json:"begin_date" binding:"required"`
	EndDate      time.Time `json:"end_date" binding:"required"`
	Purpose      string    `json:"purpose" binding:"required"`
	Visitor      string    `json:"visitor" binding:"required"`
	Branch       string    `json:"branch" binding:"required"`
	Employee     string    `json:"employee" binding:"required"`
	PassportScan string    `json:"passport_scan" binding:"required"`
	Type         string    `json:"type" binding:"required"`
	Status       string    `json:"status" binding:"required"`
}
