package grpc

import (
	"github.com/BernsteinMondy/goorder/orders/src/internal/domain"
	proto "github.com/BernsteinMondy/goorder/orders/src/pkg/grpc"
)

type OrdersGRPCAdapter struct {
	proto.UnimplementedOrdersServer
	service *domain.Service
}

func NewOrdersGRPCAdapter(service *domain.Service) *OrdersGRPCAdapter {
	return &OrdersGRPCAdapter{
		service: service,
	}
}

func (o *OrdersGRPCAdapter) CreateOrder() {
	o.CreateOrder()
}

// TODO: Finish all of this stuff
