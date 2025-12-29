package domain

import (
	"context"

	"github.com/jorgeAM/grpc-kata-order-service/pkg/errors"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/model"
)

var (
	ErrOrderInternal = errors.Define("order.internal_error")
	ErrOrderNotFound = errors.Define("order.not_found")
)

//go:generate mockgen -source=./order.go -destination=../mock/order.go -package=mock -mock_names=Repository=MockOrderRepository
type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
}

type Order struct {
	ID         model.ID
	CustomerID model.ID
	Status     OrderStatus
	Items      []OrderItem
	Timestamps model.Timestamps
}

func NewOrder(customerID model.ID) *Order {
	return &Order{
		ID:         model.GenerateUUID(),
		CustomerID: customerID,
		Status:     Pending,
		Timestamps: model.NewTimestamps(),
	}
}

func (o *Order) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
}

func (o *Order) Total() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.UnitPrice * float64(item.Quantity)
	}

	return total
}
