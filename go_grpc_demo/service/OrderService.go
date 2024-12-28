package service

import (
	"context"
	"log"

	"github.com/vignesh/grpc_demo/protogen/golang/orders"
)

type OrderService struct {
	orders.UnimplementedOrdersServer
}

func NewOrderService() OrderService {
	return OrderService{}
}
func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Println("Received an add-order request", req)

	// err := o.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, nil
}
