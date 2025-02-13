package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/infrastructure/handler"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/audit_handler"
	"go-todo-list.com/m/internal/todos/infrastructure/persistance"
)

func SetupServer(
	todoRepository domain.Repository,
	auditTodoRepository persistance.AuditRepository,
	generator domain.IDGenerator,
) http.Handler {
	server := mux.NewRouter()

	server.HandleFunc("/api/v1/todo", handler.HandleGetAllTodos(todoRepository)).Methods("GET")
	server.HandleFunc("/api/v1/todo/{id}", handler.HandleGetTodoByID(todoRepository)).Methods("GET")
	server.HandleFunc("/api/v1/todo", handler.HandleCreateTodo(todoRepository, generator)).Methods("POST")
	server.HandleFunc("/api/v1/todo/{id}", handler.HandleUpdateTodo(todoRepository)).Methods("PUT")
	server.HandleFunc("/api/v1/todo/{id}", handler.HandleDeleteTodo(todoRepository)).Methods("DELETE")

	// audit todo logs
	server.HandleFunc("/api/v1/report", audit_handler.HandleGetAllTodoAudit(auditTodoRepository)).Methods("GET")
	server.HandleFunc("/api/v1/report/{id}", audit_handler.HandleGetTodoAuditByID(auditTodoRepository)).Methods("GET")

	// cors setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	return c.Handler(server)
}
