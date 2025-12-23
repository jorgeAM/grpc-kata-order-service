package grpc

import (
	"context"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/application/command"
	orderpb "github.com/jorgeAM/grpc-kata-proto/gen/go/order/v1"
)

var _ orderpb.OrderServiceServer = (*OrderGrpcServer)(nil)

type OrderGrpcServer struct {
	createOrderApp command.CreateOrder
	*orderpb.UnimplementedOrderServiceServer
}

func (o *OrderGrpcServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	panic("no implemented")
}
