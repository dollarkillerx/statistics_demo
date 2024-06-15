package storage

import (
	"followme/internal/models"

	"sync"
)

type Storage struct {
	mu     sync.Mutex
	Orders []models.Order
}

func NewStorage() *Storage {
	return &Storage{
		Orders: make([]models.Order, 0),
	}
}

func (s *Storage) SetOrder(orders []models.Order) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Orders = orders
}

func (s *Storage) GetOrders() []models.Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Orders
}
