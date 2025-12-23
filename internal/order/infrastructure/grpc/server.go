package grpc

import (
	"context"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/application/command"
	orderpb "github.com/jorgeAM/grpc-kata-proto/gen/go/order/v1"
)

var _ orderpb.OrderServiceServer = (*OrderGrpcServer)(nil)

type OrderGrpcServer struct {
	createOrderApp *command.CreateOrder
	*orderpb.UnimplementedOrderServiceServer
}

func NewOrderGrpcServer(createOrderApp *command.CreateOrder) *OrderGrpcServer {
	return &OrderGrpcServer{
		createOrderApp:                  createOrderApp,
		UnimplementedOrderServiceServer: &orderpb.UnimplementedOrderServiceServer{},
	}
}

func (o *OrderGrpcServer) Create(ctx context.Context, request *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	items := make([]command.Item, 0, len(request.Items))

	for _, item := range request.Items {
		items = append(items, command.Item{
			ProductCode: item.ProductCode,
			Quantity:    int(item.Quantity),
			UnitPrice:   float64(item.UnitPrice),
		})
	}

	cmd := command.CreateOrderCommand{
		CustomerID: request.UserId,
		Items:      items,
	}

	orderID, err := o.createOrderApp.Exec(ctx, &cmd)
	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{
		OrderId: orderID,
	}, nil
}
