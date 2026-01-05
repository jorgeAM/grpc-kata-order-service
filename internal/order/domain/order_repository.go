package domain

import "context"

//go:generate mockgen -source=./order_repository.go -destination=../mock/order_repository.go -package=mock -mock_names=Repository=MockOrderRepository
type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
}
