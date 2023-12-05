package test

import (
	"context"
	"testing"

	"github.com/sesaquecruz/go-sql-seed/config"
	"github.com/sesaquecruz/go-sql-seed/database"
	"github.com/sesaquecruz/go-sql-seed/seed"
	"github.com/sesaquecruz/go-sql-seed/testcontainers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type user struct {
	id    string
	email string
}

func TestApplyDataToPostgres(t *testing.T) {
	ctx := context.Background()

	postgres, err := testcontainers.NewPostgresContainer(ctx)
	require.Nil(t, err)
	defer postgres.Terminate(ctx)

	cfg := config.Config{
		DbDriver:       "postgres",
		DbDsn:          postgres.Dsn(),
		MigrationsPath: "../sql/postgres/migrations/",
		DataPath:       "../sql/postgres/data/",
	}

	err = seed.ApplyAll(&cfg)
	require.Nil(t, err)

	db, err := database.Connect(cfg.DbDriver, cfg.DbDsn)
	require.Nil(t, err)

	rows, err := db.QueryContext(ctx, "SELECT id, email FROM users")
	require.Nil(t, err)
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.email)
		require.Nil(t, err)
		users = append(users, u)
	}

	require.Equal(t, 2, len(users))

	assert.Equal(t, "id_user1", users[0].id)
	assert.Equal(t, "postgres_user1@example.com", users[0].email)

	assert.Equal(t, "id_user2", users[1].id)
	assert.Equal(t, "postgres_user2@example.com", users[1].email)
}

func TestApplyDataToMySQL(t *testing.T) {
	ctx := context.Background()

	mysql, err := testcontainers.NewMySQLContainer(ctx)
	require.Nil(t, err)
	defer mysql.Terminate(ctx)

	cfg := config.Config{
		DbDriver:       "mysql",
		DbDsn:          mysql.Dsn(),
		MigrationsPath: "../sql/mysql/migrations/",
		DataPath:       "../sql/mysql/data/",
	}

	err = seed.ApplyAll(&cfg)
	require.Nil(t, err)

	db, err := database.Connect(cfg.DbDriver, cfg.DbDsn)
	require.Nil(t, err)

	rows, err := db.QueryContext(ctx, "SELECT id, email FROM users")
	require.Nil(t, err)
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.email)
		require.Nil(t, err)
		users = append(users, u)
	}

	require.Equal(t, 2, len(users))

	assert.Equal(t, "id_user1", users[0].id)
	assert.Equal(t, "mysql_user1@example.com", users[0].email)

	assert.Equal(t, "id_user2", users[1].id)
	assert.Equal(t, "mysql_user2@example.com", users[1].email)
}
