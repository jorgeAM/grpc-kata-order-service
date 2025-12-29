package port

import (
	"context"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
)

type PaymentPort interface {
	Charge(ctx context.Context, order *domain.Order) error
}
