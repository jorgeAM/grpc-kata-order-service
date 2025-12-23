package persistence

import (
	"time"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/model"
)

type postgresOrder struct {
	ID         string     `db:"id"`
	CustomerID string     `db:"customer_id"`
	Status     string     `db:"status"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" goqu:"omitnil"`
}

type postgresOrderItem struct {
	ID          string     `db:"id"`
	OrderID     string     `db:"order_id"`
	ProductCode string     `db:"product_code"`
	Quantity    int        `db:"quantity"`
	UnitPrice   float64    `db:"unit_price"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" goqu:"omitnil"`
}

func (dto postgresOrder) toDomain() (*domain.Order, error) {
	id, err := model.NewID(dto.ID)
	if err != nil {
		return nil, err
	}

	customerID, err := model.NewID(dto.CustomerID)
	if err != nil {
		return nil, err
	}

	status, err := domain.NewOrderStatus(dto.Status)
	if err != nil {
		return nil, err
	}

	return &domain.Order{
		ID:         id,
		CustomerID: customerID,
		Status:     status,
		Timestamps: model.Timestamps{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
			DeletedAt: dto.DeletedAt,
		},
	}, nil
}

func (dto postgresOrderItem) toDomain() (*domain.OrderItem, error) {
	id, err := model.NewID(dto.ID)
	if err != nil {
		return nil, err
	}

	orderID, err := model.NewID(dto.OrderID)
	if err != nil {
		return nil, err
	}

	return &domain.OrderItem{
		ID:          id,
		OrderID:     orderID,
		ProductCode: dto.ProductCode,
		Quantity:    dto.Quantity,
		UnitPrice:   dto.UnitPrice,
		Timestamps: model.Timestamps{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
			DeletedAt: dto.DeletedAt,
		},
	}, nil
}

func fromOrderDomain(entity *domain.Order) postgresOrder {
	return postgresOrder{
		ID:         entity.ID.String(),
		CustomerID: entity.CustomerID.String(),
		Status:     entity.Status.String(),
		CreatedAt:  entity.Timestamps.CreatedAt,
		UpdatedAt:  entity.Timestamps.UpdatedAt,
		DeletedAt:  entity.Timestamps.DeletedAt,
	}
}

func fromOrderItemDomain(entity *domain.OrderItem) postgresOrderItem {
	return postgresOrderItem{
		ID:          entity.ID.String(),
		OrderID:     entity.OrderID.String(),
		ProductCode: entity.ProductCode,
		Quantity:    entity.Quantity,
		UnitPrice:   entity.UnitPrice,
		CreatedAt:   entity.Timestamps.CreatedAt,
		UpdatedAt:   entity.Timestamps.UpdatedAt,
		DeletedAt:   entity.Timestamps.DeletedAt,
	}
}
