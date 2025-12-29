package config

import "github.com/jorgeAM/grpc-kata-order-service/pkg/env"

type Config struct {
	AppEnv                     string
	Port                       string
	GrpcPort                   string
	PaymentGRPCUrl             string
	PostgresHost               string
	PostgresPort               int
	PostgresDatabase           string
	PostgresUser               string
	PostgresPassword           string
	PostgresMaxIdleConnections int
	PostgresMaxOpenConnections int
}

func LoadConfig() (*Config, error) {
	return &Config{
		AppEnv:                     env.GetEnv("APP_ENV", "local"),
		Port:                       env.GetEnv("PORT", "8080"),
		GrpcPort:                   env.GetEnv("GRPC_PORT", "9090"),
		PaymentGRPCUrl:             env.GetEnv("PAYMENT_GRPC_URL", "localhost:9091"),
		PostgresHost:               env.GetEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:               env.GetEnv("POSTGRES_PORT", 5432),
		PostgresDatabase:           env.GetEnv("POSTGRES_DB", "db"),
		PostgresUser:               env.GetEnv("POSTGRES_USER", "admin"),
		PostgresPassword:           env.GetEnv("POSTGRES_PASSWORD", "passwd123"),
		PostgresMaxIdleConnections: env.GetEnv("POSTGRES_MAX_IDLE_CONNECTIONS", 10),
		PostgresMaxOpenConnections: env.GetEnv("POSTGRES_MAX_OPEN_CONNECTIONS", 30),
	}, nil
}
