package main

import (
	"github.com/gorilla/mux"

	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/infrastructure/handler"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/audit_handler"
	"go-todo-list.com/m/internal/todos/infrastructure/persistance"
)

func SetupServer(
	todoRepository domain.Repository,
	auditTodoRepository persistance.AuditRepository,
	generator domain.IDGenerator,
) *mux.Router {
	server := mux.NewRouter()

	server.HandleFunc("/api/v1/todo", handler.HandleGetAllTodos(todoRepository)).Methods("GET")
	server.HandleFunc("/api/v1/todo/{id}", handler.HandleGetTodoByID(todoRepository)).Methods("GET")
	server.HandleFunc("/api/v1/todo", handler.HandleCreateTodo(todoRepository, generator)).Methods("POST")
	server.HandleFunc("/api/v1/todo/{id}", handler.HandleUpdateTodo(todoRepository)).Methods("PUT")
	server.HandleFunc("/api/v1/todo/{id}", handler.HandleDeleteTodo(todoRepository)).Methods("DELETE")

	// audit todo logs
	server.HandleFunc("/api/v1/todo/report", audit_handler.HandleGetAllTodoAudit(auditTodoRepository)).Methods("GET")
	server.HandleFunc("/api/v1/todo/report/{id}", audit_handler.HandleGetTodoAuditByID(auditTodoRepository)).Methods("GET")

	return server
}
