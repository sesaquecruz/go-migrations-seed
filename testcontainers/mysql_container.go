package testcontainers

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

type MySQLContainer struct {
	container *mysql.MySQLContainer
	dsn       string
}

func NewMySQLContainer(ctx context.Context) (*MySQLContainer, error) {
	username := "mysql"
	password := "mysql"
	database := "test_db"

	container, err := mysql.RunContainer(ctx,
		testcontainers.WithImage("mysql:8.0-debian"),
		mysql.WithUsername(username),
		mysql.WithPassword(password),
		mysql.WithDatabase(database),
	)
	if err != nil {
		return nil, err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := container.MappedPort(ctx, "3306/tcp")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false",
		username, password, host, port.Port(), database,
	)

	return &MySQLContainer{
		container: container,
		dsn:       dsn,
	}, nil
}

func (c *MySQLContainer) Dsn() string {
	return c.dsn
}

func (c *MySQLContainer) Terminate(ctx context.Context) error {
	return c.container.Terminate(ctx)
}
