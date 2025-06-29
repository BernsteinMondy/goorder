package domain

type OrderStatus uint8

const (
	OrderStatusCreated OrderStatus = iota + 1
	OrderStatusPayed
	OrderStatusSuccess
	OrderStatusCancelled
)

type ProductSize uint8

const (
	ProductSizeSmall = iota + 1
	ProductSizeMedium
	ProductSizeLarge
	ProductSizeExtraLarge
)
