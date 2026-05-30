package service

import (
	model "order-service/internal/entity"
	"order-service/internal/repository"
)

type OrderService struct{
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService{
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(order model.Order) error {
	err := s.repo.CreateOrder(order)
	return err
}

func (s *OrderService) GetOrderByID(ID string) (model.Order, error) {
	order,err := s.repo.GetOrderByID(ID)
	if err != nil {
		return  model.Order{},err
	}
	return order,nil
}


func (s *OrderService) GetAllOrders() ([]model.Order, error) {
	order,err := s.repo.GetAllOrders()
	if err != nil {
		return  []model.Order{},err
	}
	return order,nil
}