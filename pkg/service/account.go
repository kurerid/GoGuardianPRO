package service

import (
	"GuardianPRO/models"
	"GuardianPRO/pkg/repository"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) IsRegistered(email string) (bool, error) {
	return s.repo.IsRegistered(email)
}

func (s *AccountService) GetList() (*models.AccountGetList, error) {
	return s.repo.GetList()
}
