package persistence

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/jorgeAM/grpc-kata-order-service/internal/order/domain"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/model"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresOrderRepositoryTestSuite struct {
	suite.Suite
	postgresDockerContainer *testcontainers.DockerContainer
	db                      *sqlx.DB
}

func (p *PostgresOrderRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	postgres, err := testcontainers.Run(
		ctx,
		"postgres:latest",
		testcontainers.WithExposedPorts("5432/tcp"),
		testcontainers.WithEnv(map[string]string{
			"POSTGRES_USER":     "admin",
			"POSTGRES_PASSWORD": "passwd123",
			"POSTGRES_DB":       "mydb",
		}),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("5432/tcp"),
			wait.ForLog("database system is ready to accept connections"),
		),
	)

	p.postgresDockerContainer = postgres
	require.NoError(p.T(), err)

	host, err := postgres.Host(ctx)
	require.NoError(p.T(), err)

	port, err := postgres.MappedPort(ctx, "5432")
	require.NoError(p.T(), err)

	dsn := fmt.Sprintf("host=%s port=%s user=admin password=passwd123 dbname=mydb sslmode=disable", host, port.Port())
	p.db, err = sqlx.Connect("postgres", dsn)
	require.NoError(p.T(), err)

	p.setupDatabase()
}

func (p *PostgresOrderRepositoryTestSuite) TearDownSuite() {
	if p.db != nil {
		p.db.Close()
	}

	testcontainers.CleanupContainer(p.T(), p.postgresDockerContainer)
}

func (p *PostgresOrderRepositoryTestSuite) setupDatabase() {
	migrationPath := "../../../../database/migration"

	entries, err := os.ReadDir(migrationPath)
	require.NoError(p.T(), err)

	var upFiles []string
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".up.sql") {
			upFiles = append(upFiles, entry.Name())
		}
	}

	sort.Strings(upFiles)

	for _, filename := range upFiles {
		content, err := os.ReadFile(filepath.Join(migrationPath, filename))
		require.NoError(p.T(), err)

		p.T().Logf("Executing migration: %s", filename)
		_, err = p.db.Exec(string(content))
		require.NoError(p.T(), err, "failed to execute migration: %s", filename)
	}
}

func (p *PostgresOrderRepositoryTestSuite) TestShouldSaveOrder() {
	ctx := context.Background()
	orderRepository := NewPostgresOrderRepository(p.db)

	order := domain.NewOrder(model.GenerateUUID())
	order.AddItem(domain.NewOrderItem(
		order.ID,
		"product-1",
		2,
		9.99,
	))

	err := orderRepository.Save(ctx, order)
	p.Nil(err, "err should be nil")
}

func TestPostgresOrderRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresOrderRepositoryTestSuite))
}
