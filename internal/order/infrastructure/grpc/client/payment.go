package client

import (
	"context"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/application/port"
	"github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
	paymentpb "github.com/jorgeAM/grpc-kata-proto/gen/go/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ port.PaymentPort = (*PaymentGRPCClient)(nil)

type PaymentGRPCClient struct {
	paymentClient paymentpb.PaymentServiceClient
	conn          *grpc.ClientConn
}

func NewPaymentGRPCClient(paymentServiceURL string) (*PaymentGRPCClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(paymentServiceURL, opts...)
	if err != nil {
		return nil, err
	}

	paymentClient := paymentpb.NewPaymentServiceClient(conn)

	return &PaymentGRPCClient{
		paymentClient: paymentClient,
		conn:          conn,
	}, nil
}

func (p *PaymentGRPCClient) Charge(ctx context.Context, order *domain.Order) error {
	_, err := p.paymentClient.Create(ctx, &paymentpb.CreatePaymentRequest{
		OrderId:    order.ID.String(),
		UserId:     order.CustomerID.String(),
		TotalPrice: float32(order.Total()),
	})
	if err != nil {
		return err
	}

	return nil
}
func (p *PaymentGRPCClient) Close() error {
	return p.conn.Close()
}
