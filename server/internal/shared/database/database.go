package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	User         string
	Password     string
	DatabaseName string
	Port         int
}

func GenerateDatabaseConnection(params DatabaseConnection) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%d/%s?sslmode=disable", params.User, params.Password, params.Port, params.DatabaseName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
