package service

import (
	"errors"

	"auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(email string, password string) (int, error) {

	if email == "" || password == "" {
		return 0, errors.New("email or password is empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.CreateUser(email, string(hashedPassword))
	if err != nil {
		return 0, err
	}

	return id, nil
}