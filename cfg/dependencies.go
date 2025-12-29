package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jorgeAM/grpc-kata-order-service/internal/order/application/port"
	orderDomain "github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
	"github.com/jorgeAM/grpc-kata-order-service/internal/order/infrastructure/grpc/client"
	orderPersistence "github.com/jorgeAM/grpc-kata-order-service/internal/order/infrastructure/persistence"
	_ "github.com/lib/pq"
)

type Dependencies struct {
	OrderRepository orderDomain.OrderRepository
	PaymentPort     port.PaymentPort
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

	paymentGRPCClient, err := client.NewPaymentGRPCClient(cfg.PaymentGRPCUrl)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		OrderRepository: postgresOrderRepository,
		PaymentPort:     paymentGRPCClient,
	}, nil
}
