package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         int
}

func GenerateDatabaseConnection(params DatabaseConnection) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", params.Host, params.Port, params.User, params.Password, params.DatabaseName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
