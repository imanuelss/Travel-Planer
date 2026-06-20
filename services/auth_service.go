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

// this is fu0nction for login logic, it will check if the email exists and if the password is correct, then generate a JWT token
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !helpers.CheckPassword(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
