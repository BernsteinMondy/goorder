package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/BernsteinMondy/goorder/orders/src/internal/domain"
	"github.com/google/uuid"
	"strings"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (err error) {
	const (
		orderQuery      = `INSERT INTO order.orders (row_id,row_created_at,id,created_at,user_id,status) VALUES ($1,$2,$3,$4,$5,$6)`
		orderItemsQuery = `INSERT INTO order.items (id,order_id,product_id,quantity) VALUES %s`
	)

	tx, err := o.db.BeginTx(ctx, nil)
	defer func() {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			if !errors.Is(err, sql.ErrTxDone) {
				err = errors.Join(err, fmt.Errorf("rollback sql tx: %w", rollBackErr))
			}
		}
	}()
	if err != nil {
		return fmt.Errorf("being sql tx: %w", err)
	}

	statusSQL := orderStatusToSQLEnum(order.Status)

	_, err = tx.ExecContext(ctx, orderQuery, uuid.New(), time.Now(), order.ID, order.CreatedAt, order.UserID, statusSQL)
	if err != nil {
		return fmt.Errorf("run sql query: %w", err)
	}

	if len(order.Items) > 0 {
		var (
			valueStrings = make([]string, 0, len(order.Items))
			valueArgs    = make([]interface{}, 0, len(order.Items)*4)
			argPos       = 1
		)

		for _, item := range order.Items {
			valueStrings = append(valueStrings,
				fmt.Sprintf("($%d, $%d, $%d, $%d)",
					argPos, argPos+1, argPos+2, argPos+3))

			valueArgs = append(valueArgs,
				uuid.New(),
				order.ID,
				item.ProductID,
				item.Quantity,
			)

			argPos += 4 // increase argument position by amount of arguments (4)
		}

		_, err = tx.ExecContext(
			ctx,
			fmt.Sprintf(orderItemsQuery, strings.Join(valueStrings, ",")),
			valueArgs...,
		)
		if err != nil {
			return fmt.Errorf("run sql query with batch insert: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit sql tx: %w", err)
	}

	return nil
}

func (o *OrderRepository) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	const (
		orderQuery = `
			SELECT 
	    	id,
	    	created_at,
	    	user_id,
	    	status 
			FROM order.orders 
			WHERE id=$1 
			ORDER BY row_created_at DESC 
			LIMIT 1
			`
		orderItemQuery = `
			SELECT
			id,
            order_id,
            product_id,
			quantity
			FROM order.items
			WHERE order_id=$1
			`
	)

	// TODO: finish

	var (
		ret = domain.Order{
			ID: orderID,
		}
	)

	err := o.db.QueryRowContext(ctx, orderQuery, orderID).Scan(&ret.ID, &ret.CreatedAt, &ret.UserID, &ret.Status)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &ret, nil
}

func (o *OrderRepository) GetOrdersByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepository) UpdateOrderStatus(ctx context.Context, order *domain.Order) error {
	const query = `INSERT INTO order.orders (row_id,row_created_at,id,created_at,user_id,status) VALUES ($1,$2,$3,$4,$5,$6)`

	// TODO: finish

	statusSQL := orderStatusToSQLEnum(order.Status)
	_, err := o.db.ExecContext(ctx, query, uuid.New(), time.Now(), order.ID, order.CreatedAt, order.UserID, statusSQL)
	if err != nil {
		return fmt.Errorf("run sql query: %w", err)
	}

	return nil
}
