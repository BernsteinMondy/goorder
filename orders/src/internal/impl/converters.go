package impl

import (
	"fmt"
	"github.com/BernsteinMondy/goorder/orders/src/internal/domain"
)

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
