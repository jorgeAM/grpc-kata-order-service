package domain

import (
	"strings"

	"github.com/jorgeAM/grpc-kata-order-service/pkg/errors"
)

var (
	ErrInvalidOrderStatus = errors.Define("order.invalid_status")
)

type OrderStatus string

const (
	Pending OrderStatus = "PENDING"
)

var allowedOrderStatus = map[string]OrderStatus{
	Pending.String(): Pending,
}

func NewOrderStatus(t string) (OrderStatus, error) {
	if status, ok := allowedOrderStatus[strings.ToUpper(t)]; ok {
		return status, nil
	}

	return "", errors.New(
		ErrInvalidOrderStatus,
		"invalid order status",
		errors.WithMetadata("order_status", t),
	)
}

func (o OrderStatus) String() string {
	return string(o)
}
