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

func (s *AuthService) SignUp(input *models.Account) (*string, error) {
	return s.repo.SignUp(input)
}

func (s *AuthService) SignIn(input *models.Account) (*string, error) {
	return s.repo.SignIn(input)
}
