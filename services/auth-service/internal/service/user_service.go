package service

import (
	"errors"

	"auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
	jwtSvc *JWTService
}

func NewUserService(repo *repository.UserRepository,jwtSvc *JWTService) *UserService {
	return &UserService{repo: repo,
	jwtSvc: jwtSvc,}
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

func (s *UserService) Login(email string, password string) (string, error) {
	dbUser, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := s.jwtSvc.GenerateToken(int(dbUser.ID), email)
	if err != nil {
		return "", err
	}
	return token, nil
}
