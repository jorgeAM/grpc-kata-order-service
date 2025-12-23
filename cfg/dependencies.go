package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	orderDomain "github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
	orderPersistence "github.com/jorgeAM/grpc-kata-order-service/internal/order/infrastructure/persistence"
	_ "github.com/lib/pq"
)

type Dependencies struct {
	OrderRepository orderDomain.OrderRepository
}

func BuildDependencies(cfg *Config) (*Dependencies, error) {
	postgresClient, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		return nil, err
	}

	postgresOrderRepository := orderPersistence.NewPostgresOrderRepository(postgresClient)

	return &Dependencies{
		OrderRepository: postgresOrderRepository,
	}, nil
}
