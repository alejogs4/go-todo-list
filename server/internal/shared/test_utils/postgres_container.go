package test_utils

import (
	"context"
	"database/sql"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type postgresDBContainer struct {
	DB        *sql.DB
	Container *postgres.PostgresContainer
}

func GenerateTestPostgresDB(ctx context.Context, initScript, dbName string) (postgresDBContainer, error) {
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("304971447450.dkr.ecr.eu-west-3.amazonaws.com/postgres:13"),
		postgres.WithInitScripts(filepath.Join("testdata", "testcontainers", initScript)),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbName),
		postgres.WithPassword(dbName),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(30*time.Second), wait.ForExposedPort()),
	)
	if err != nil {
		return postgresDBContainer{}, err
	}

	connString, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return postgresDBContainer{}, err

	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return postgresDBContainer{}, err
	}

	return postgresDBContainer{
		DB:        db,
		Container: postgresContainer,
	}, nil
}
