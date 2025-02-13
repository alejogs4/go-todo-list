package response

import (
	"time"

	"go-todo-list.com/m/internal/todos/domain"
)

type TodoResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Completed   bool    `json:"completed"`
	CreatedAt   string  `json:"created_at"`
	CompletedAt *string `json:"completed_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateTodoResponse struct {
	ID string `json:"id"`
}

func FromDomainTodosToResponse(todos []domain.Todo) []TodoResponse {
	response := make([]TodoResponse, 0, len(todos))

	for _, todo := range todos {
		response = append(response, FromTodoToResponse(todo))
	}

	return response
}

func FromTodoToResponse(todo domain.Todo) TodoResponse {
	return TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt.Format(time.DateTime),
		CompletedAt: formatCompletedAt(todo.CompletedAt),
		UpdatedAt:   todo.UpdatedAt.Format(time.DateTime),
	}
}

func formatCompletedAt(completedAt *time.Time) *string {
	if completedAt == nil {
		return nil
	}

	completedAtString := completedAt.Format(time.DateTime)

	return &completedAtString
}
