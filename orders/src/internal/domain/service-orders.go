package domain

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	// GetOrderByID must return ErrRepoNotFound if no order was found by the provided ID.
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*Order, error)
	GetOrdersByUserID(ctx context.Context, userID uuid.UUID) ([]Order, error)
	// UpdateOrderStatus inserts new Order object state with new OrderStatus.
	UpdateOrderStatus(ctx context.Context, order *Order) error
}

func (s *Service) CreateOrder(ctx context.Context, order *Order) (uuid.UUID, error) {
	err := s.orderRepository.CreateOrder(ctx, order)
	if err != nil {
		return uuid.Nil, fmt.Errorf("order repository: create order: %w", err)
	}

	return order.ID, nil
}
