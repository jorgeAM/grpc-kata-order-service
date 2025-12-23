package domain

import "github.com/jorgeAM/grpc-kata-order-service/pkg/model"

type OrderItem struct {
	ID          model.ID
	OrderID     model.ID
	ProductCode string
	Quantity    int
	UnitPrice   float64
	Timestamps  model.Timestamps
}

func NewOrderItem(orderID model.ID, productCode string, quantity int, unitPrice float64) OrderItem {
	return OrderItem{
		ID:          model.GenerateUUID(),
		OrderID:     orderID,
		ProductCode: productCode,
		Quantity:    quantity,
		UnitPrice:   unitPrice,
		Timestamps:  model.NewTimestamps(),
	}
}
