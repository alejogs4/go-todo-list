package handler

import (
	"errors"
	"net/http"

	"go-todo-list.com/m/internal/shared/api"
	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/response"
	"go-todo-list.com/m/internal/todos/use_cases"
)

func HandleGetAllTodos(todoRepository domain.Repository) http.HandlerFunc {
	getAllTodosUseCase := use_cases.NewGetTodosUseCase(todoRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		completed := api.GetQueryParam(r, "completed")
		if completed != "" && completed != "true" && completed != "false" {
			api.InvalidRequest(w, "Invalid query param it must be true or false")
			return
		}

		filters := domain.TodoFilters{}
		if completed != "" {
			isCompleted := completed == "true"
			filters.Completed = &isCompleted
		}

		ctx := r.Context()
		todos, err := getAllTodosUseCase.Execute(ctx, filters)
		if err != nil {
			if errors.Is(err, domain.ErrTodoNotFound) && filters.Completed != nil {
				api.NotFound(w, "Todos not found")
				return
			}

			if errors.Is(err, domain.ErrTodoNotFound) && filters.Completed == nil {
				api.Success(w, []response.TodoResponse{})
				return
			}

			api.InternalServerError(w, err.Error())
			return
		}

		api.Success(w, response.FromDomainTodosToResponse(todos))
	}
}
