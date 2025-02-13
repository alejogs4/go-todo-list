package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go-todo-list.com/m/internal"
	"go-todo-list.com/m/internal/shared/database"
	"go-todo-list.com/m/internal/shared/logger"
	"go-todo-list.com/m/internal/shared/uuid"
	"go-todo-list.com/m/internal/todos/infrastructure/persistance"
)

func main() {
	databaseConnection, err := database.GenerateDatabaseConnection(getDatabaseParamsFromEnv())
	if err != nil {
		panic(err)
	}

	output, err := logger.NewLoggerWriter("todo-audit.log")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	appLogger := log.New(output, "todo-audit", log.LstdFlags)

	auditRepository := persistance.NewPostgresAuditRepository(databaseConnection)
	todoRepository := persistance.NewTodosPostgresRepository(databaseConnection)
	todoAuditRepository := persistance.NewTodoAuditRepositoryDecorator(todoRepository, auditRepository, appLogger)

	server := internal.SetupServer(todoAuditRepository, auditRepository, uuid.UUIDGenerator{})

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), server)
	if err != nil {
		panic(err)
	}
}

func getDatabaseParamsFromEnv() database.DatabaseConnection {
	port := 5432
	if os.Getenv("DB_PORT") != "" {
		if portInt, err := strconv.Atoi(os.Getenv("DB_PORT")); err == nil {
			port = portInt
		}
	}

	return database.DatabaseConnection{
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:         port,
		Host:         os.Getenv("DB_HOST"),
	}
}
