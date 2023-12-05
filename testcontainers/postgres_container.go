package testcontainers

import (
	"context"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	container *postgres.PostgresContainer
	dsn       string
}

func NewPostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	username := "postgres"
	password := "postgres"
	database := "test_db"

	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16.0-alpine"),
		postgres.WithUsername(username),
		postgres.WithPassword(password),
		postgres.WithDatabase(database),
		testcontainers.WithWaitStrategy(
			wait.
				ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		return nil, err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := container.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@%s:%s/%s?sslmode=disable",
		username, password, host, port.Port(), database,
	)

	return &PostgresContainer{
		container: container,
		dsn:       dsn,
	}, nil
}

func (c *PostgresContainer) Dsn() string {
	return c.dsn
}

func (c *PostgresContainer) Terminate(ctx context.Context) error {
	return c.container.Terminate(ctx)
}
