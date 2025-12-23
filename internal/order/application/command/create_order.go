package command

import (
	"context"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/model"
)

type Item struct {
	ProductCode string  `json:"product_code"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

type CreateOrderCommand struct {
	CustomerID string `json:"customer_id"`
	Items      []Item `json:"items"`
}

type CreateOrder struct {
	orderRepository domain.OrderRepository
}

func NewCreateOrder(orderRepository domain.OrderRepository) *CreateOrder {
	return &CreateOrder{
		orderRepository: orderRepository,
	}
}

func (c *CreateOrder) Exec(ctx context.Context, cmd *CreateOrderCommand) error {
	customerID, err := model.NewID(cmd.CustomerID)
	if err != nil {
		return err
	}

	order := domain.NewOrder(customerID)

	for _, item := range cmd.Items {
		order.AddItem(domain.NewOrderItem(order.ID, item.ProductCode, item.Quantity, item.UnitPrice))
	}

	return c.orderRepository.Save(ctx, order)
}
