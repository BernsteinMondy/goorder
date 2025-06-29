package domain

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *Product) error
	// GetProductByID must return ErrRepoNotFound if no product was found by the provided ID.
	GetProductByID(ctx context.Context, productID uuid.UUID) (*Product, error)
	DeleteProductByID(ctx context.Context, productID uuid.UUID) error
}

func (s *Service) CreateProduct(ctx context.Context, product *Product) error {
	err := s.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return fmt.Errorf("product repository: create product: %w", err)
	}

	return nil
}
