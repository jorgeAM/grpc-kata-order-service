package port

import (
	"context"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
)

//go:generate mockgen -source=./payment.go -destination=../../mock/payment_port.go -package=mock -mock_names=Repository=MockPaymentClient
type PaymentPort interface {
	Charge(ctx context.Context, order *domain.Order) error
}
