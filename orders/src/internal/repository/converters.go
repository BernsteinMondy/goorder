package repository

import (
	"errors"
	"fmt"
	"github.com/BernsteinMondy/goorder/orders/src/internal/domain"
)

type OrderStatusScanner struct {
	Status domain.OrderStatus
}

func (o *OrderStatusScanner) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src is nil")
	}

	switch v := src.(type) {
	case string:
		return o.scanString(v)
	case []byte:
		return o.scanString(string(v))
	default:
		return fmt.Errorf("cannot convert type %T to domain.OrderStatus", src)
	}

}

func (o *OrderStatusScanner) scanString(v string) error {
	switch v {
	case "created":
		o.Status = domain.OrderStatusCreated
		return nil
	case "payed":
		o.Status = domain.OrderStatusPayed
		return nil
	case "cancelled":
		o.Status = domain.OrderStatusCancelled
		return nil
	case "success":
		o.Status = domain.OrderStatusSuccess
		return nil
	default:
		return fmt.Errorf("unknown enum value %s", v)
	}
}

func orderStatusToSQLEnum(status domain.OrderStatus) string {
	switch status {
	case domain.OrderStatusCreated:
		return "created"
	case domain.OrderStatusCancelled:
		return "cancelled"
	case domain.OrderStatusPayed:
		return "payed"
	case domain.OrderStatusSuccess:
		return "success"
	default:
		panic(fmt.Sprintf("unknown domain.OrderStatus: (%v)", status))
	}

}

type ProductSizeScanner struct {
	Size domain.ProductSize
}

func (p *ProductSizeScanner) Scan(src interface{}) error {
	if src == nil {
		return errors.New("src is nil")
	}

	switch v := src.(type) {
	case string:
		return p.scanString(v)
	case []byte:
		return p.scanString(string(v))
	default:
		return fmt.Errorf("cannot convert type %T to domain.ProductSize", src)
	}

}

func (p *ProductSizeScanner) scanString(v string) error {
	switch v {
	case "small":
		p.Size = domain.ProductSizeSmall
		return nil
	case "medium":
		p.Size = domain.ProductSizeMedium
		return nil
	case "large":
		p.Size = domain.ProductSizeLarge
		return nil
	case "extra_large":
		p.Size = domain.ProductSizeExtraLarge
		return nil
	default:
		return fmt.Errorf("unknown enum value %s", v)
	}
}

func productSizeToSQLEnum(size domain.ProductSize) string {
	switch size {
	case domain.ProductSizeLarge:
		return "large"
	case domain.ProductSizeExtraLarge:
		return "extra_large"
	case domain.ProductSizeSmall:
		return "small"
	case domain.ProductSizeMedium:
		return "medium"
	default:
		panic(fmt.Sprintf("unknown domain.ProductSize: (%v)", size))
	}
}
