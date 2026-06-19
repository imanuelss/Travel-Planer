package services

import (
	"errors"

	"individual-project/helpers"
	"individual-project/models"
	"individual-project/repositories"
)

// service struct
type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{repo}
}

// register business logic
func (s *AuthService) Register(name, email, password string) (*models.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("all fields required")
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
