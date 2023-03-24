package service

import (
	"GuardianPRO/models"
	"GuardianPRO/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(account *models.Account) (*string, error) {
	return s.repo.SignUp(account)
}

func (s *AuthService) SignIn(account *models.Account) (*string, error) {
	return s.repo.SignIn(account)
}
