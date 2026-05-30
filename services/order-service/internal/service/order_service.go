package service

import (
	"errors"
	model "order-service/internal/entity"
	"order-service/internal/grpc/client"
	"order-service/internal/repository"
)

type OrderService struct {
	repo       *repository.OrderRepository
	authClient *client.AuthClient
}

func NewOrderService(repo *repository.OrderRepository, authClient *client.AuthClient) *OrderService {
	return &OrderService{
		repo:       repo,
		authClient: authClient,
	}
}


func (s *OrderService) CreateOrder(order model.Order) error {
	user,err := s.authClient.GetUserByID(int32(order.UserID))
	if err != nil {
		return errors.New("auth service error or user not found")
	}

	if user.Id == 0 {
		return errors.New("invalid user")
	}
	err = s.repo.CreateOrder(order)
	return err
}

func (s *OrderService) GetOrderByID(ID string,UserID int) (model.Order, error) {
	user,err := s.authClient.GetUserByID(int32(UserID))
	if err != nil {
		return model.Order{}, errors.New("auth service error or user not found")
	}

	if user.Id == 0 {
		return model.Order{}, errors.New("invalid user")
	}
	order,err := s.repo.GetOrderByID(ID)
	if err != nil {
		return  model.Order{},err
	}
	return order,nil
}


func (s *OrderService) GetAllOrders(UserID int) ([]model.Order, error) {
	user,err := s.authClient.GetUserByID(int32(UserID))
	if err != nil {
		return []model.Order{}, errors.New("auth service error or user not found")
	}

	if user.Id == 0 {
		return []model.Order{}, errors.New("invalid user")
	}
	order,err := s.repo.GetAllOrders()
	if err != nil {
		return  []model.Order{},err
	}
	return order,nil
}