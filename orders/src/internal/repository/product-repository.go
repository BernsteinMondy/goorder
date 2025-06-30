package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/BernsteinMondy/goorder/orders/src/internal/domain"
	"github.com/google/uuid"
)

type ProductRepository struct {
	db *sql.DB
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	const query = `INSERT INTO product.products (id,name,price,size) VALUES ($1,$2,$3,$4)`

	sizeSQL := productSizeToSQLEnum(product.Size)
	_, err := p.db.ExecContext(ctx, query, product.ID, product.Name, product.Price, sizeSQL)
	if err != nil {
		return fmt.Errorf("run sql query: %w", err)
	}

	return nil
}

func (p *ProductRepository) GetProductByID(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	const query = `SELECT id,name,price,size FROM product.products WHERE id = $1`

	var (
		ret domain.Product
		psc ProductSizeScanner
	)

	err := p.db.QueryRowContext(ctx, query, productID).Scan(&ret.ID, &ret.Name, &ret.Price, &psc)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrRepoNotFound
		}
		return nil, fmt.Errorf("query row: %w", err)
	}

	ret.Size = psc.Size

	return &ret, nil
}

func (p *ProductRepository) DeleteProductByID(ctx context.Context, productID uuid.UUID) error {
	const query = `DELETE FROM product.products WHERE id = $1`

	_, err := p.db.ExecContext(ctx, query, productID)
	if err != nil {
		return fmt.Errorf("run sql query: %w", err)
	}

	return nil
}
