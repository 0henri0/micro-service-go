package services

import (
	"micro-service-go/internal/models"
	"micro-service-go/internal/repositories"
)

type OrderService interface {
	GetOrders() ([]models.Order, error)
}

type orderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (u orderService) GetOrders() ([]models.Order, error) {
	return u.orderRepo.GetOrders()
}
