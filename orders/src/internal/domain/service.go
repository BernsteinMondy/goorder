package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type (
	Order struct {
		ID        uuid.UUID // Immutable, unique
		CreatedAt time.Time
		UserID    uuid.UUID
		Items     []OrderItem
		Status    OrderStatus
	}
	OrderItem struct {
		ID        uuid.UUID // Immutable, unique
		OrderID   uuid.UUID
		ProductID uuid.UUID
		Quantity  int
	}
	Product struct {
		ID    uuid.UUID // Immutable, unique
		Name  string
		Price decimal.Decimal
		Size  ProductSize
	}
)

type Service struct {
	orderRepository   OrderRepository
	productRepository ProductRepository
}

func NewService(
	orderRepository OrderRepository,
	productRepository ProductRepository,
) *Service {
	return &Service{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}
