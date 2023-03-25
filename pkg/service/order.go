package service

import (
	"GuardianPRO/models"
	"GuardianPRO/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreatePrivate(input *models.Order) error {
	return s.repo.CreatePrivate(input)
}
