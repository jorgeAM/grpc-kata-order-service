package persistence

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/errors"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var _ domain.OrderRepository = (*PostgresOrderRepository)(nil)

type PostgresOrderRepository struct {
	db     *sqlx.DB
	schema string
}

func NewPostgresOrderRepository(db *sqlx.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{
		db:     db,
		schema: "order_schema",
	}
}

func (p *PostgresOrderRepository) Save(ctx context.Context, order *domain.Order) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer tx.Rollback()

	orderDTO := fromOrderDomain(order)
	orderDS := goqu.
		Dialect("postgres").
		Insert(fmt.Sprintf("%s.orders", p.schema)).
		Rows(orderDTO)

	orderSQL, orderArgs, err := orderDS.Prepared(true).ToSQL()
	if err != nil {
		return errors.Wrap(domain.ErrOrderInternal, err, "we got a problem generating order sql statement")
	}

	_, err = tx.ExecContext(ctx, orderSQL, orderArgs...)
	if err != nil {
		return errors.Wrap(domain.ErrOrderInternal, err, "an error occurred while saving order")
	}

	if len(order.Items) > 0 {
		itemDTOs := make([]interface{}, 0, len(order.Items))
		for _, item := range order.Items {
			itemDTOs = append(itemDTOs, fromOrderItemDomain(&item))
		}

		itemsDS := goqu.
			Dialect("postgres").
			Insert(fmt.Sprintf("%s.order_items", p.schema)).
			Rows(itemDTOs...)

		itemsSQL, itemsArgs, err := itemsDS.Prepared(true).ToSQL()
		if err != nil {
			return errors.Wrap(domain.ErrOrderInternal, err, "we got a problem generating order items sql statement")
		}

		_, err = tx.ExecContext(ctx, itemsSQL, itemsArgs...)
		if err != nil {
			return errors.Wrap(domain.ErrOrderInternal, err, "an error occurred while saving order items")
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(domain.ErrOrderInternal, err, "failed to commit transaction")
	}

	return nil
}
