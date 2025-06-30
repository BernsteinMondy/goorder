package domain

type OrderStatus uint8

const (
	OrderStatusCreated OrderStatus = iota
	OrderStatusPayed
	OrderStatusSuccess
	OrderStatusCancelled
)

type ProductSize uint8

const (
	ProductSizeSmall = iota
	ProductSizeMedium
	ProductSizeLarge
	ProductSizeExtraLarge
)
