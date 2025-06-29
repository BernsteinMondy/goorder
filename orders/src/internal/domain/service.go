package domain

import (
	"context"
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

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) (uuid.UUID, error)
	// GetOrderByID must return ErrRepoNotFound if no order was found by the provided ID.
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*Order, error)
	GetOrdersByUserID(ctx context.Context, userID uuid.UUID) ([]Order, error)
	// UpdateOrderStatus inserts new Order object state with new OrderStatus.
	UpdateOrderStatus(ctx context.Context, order *Order) error
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *Product) error
	// GetProductByID must return ErrRepoNotFound if no product was found by the provided ID.
	GetProductByID(ctx context.Context, productID uuid.UUID) (*Product, error)
	DeleteProductByID(ctx context.Context, productID uuid.UUID) error
}

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
