package service

import (
	"encoding/json"
	"errors"
	"fmt"
	model "order-service/internal/entity"
	"order-service/internal/grpc/client"
	"order-service/internal/repository"
	"order-service/rabbitmq"

	"github.com/google/uuid"
)

type OrderService struct {
	repo       *repository.OrderRepository
	authClient *client.AuthClient
	rabbit     *rabbitmq.RabbitMQ
}

type Event struct {
	ID         string `json:"id"`
	Body       string `json:"body"`
	RetryCount int    `json:"retry_count"`
}

func NewOrderService(repo *repository.OrderRepository, authClient *client.AuthClient, rabbit *rabbitmq.RabbitMQ) *OrderService {
	return &OrderService{
		repo:       repo,
		authClient: authClient,
		rabbit:     rabbit,
	}
}

func (s *OrderService) CreateOrder(order model.Order) error {
	user, err := s.authClient.GetUserByID(int32(order.UserID))
	if err != nil {
		return errors.New("auth service error or user not found")
	}

	if user.Id == 0 {
		return errors.New("invalid user")
	}
	err = s.repo.CreateOrder(order)
	if err != nil {
		return err
	}
	event := Event{
		ID:         uuid.New().String(),
		Body:       fmt.Sprintf("order created by user %d", order.UserID),
		RetryCount: 0,
	}
	
	data, _ := json.Marshal(event)
	
	s.rabbit.Publish(string(data))
	// for fake fail massage
	// err = s.rabbit.Publish(
	// 	"fail"	)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *OrderService) GetOrderByID(ID string, UserID int) (model.Order, error) {
	user, err := s.authClient.GetUserByID(int32(UserID))
	if err != nil {
		return model.Order{}, errors.New("auth service error or user not found")
	}

	if user.Id == 0 {
		return model.Order{}, errors.New("invalid user")
	}
	order, err := s.repo.GetOrderByID(ID)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (s *OrderService) GetAllOrders(UserID int) ([]model.Order, error) {
	user, err := s.authClient.GetUserByID(int32(UserID))
	if err != nil {
		return []model.Order{}, errors.New("auth service error or user not found")
	}

	if user.Id == 0 {
		return []model.Order{}, errors.New("invalid user")
	}
	order, err := s.repo.GetAllOrders()
	if err != nil {
		return []model.Order{}, err
	}
	return order, nil
}
