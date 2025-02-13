package handler

import (
	"errors"
	"fmt"
	"net/http"

	"go-todo-list.com/m/internal/shared/api"
	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/response"
	"go-todo-list.com/m/internal/todos/use_cases"
)

func HandleGetTodoByID(todoRepository domain.Repository) http.HandlerFunc {
	getTodoByIDUseCase := use_cases.NewGetTodosByIDUseCase(todoRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		todoID, err := api.GetPathParam(r, "id")
		if err != nil {
			api.InvalidRequest(w, "Missing post id")
			return
		}

		ctx := r.Context()
		todo, err := getTodoByIDUseCase.Execute(ctx, domain.TodoFilters{ID: &todoID})
		if err != nil {
			if errors.Is(err, domain.ErrTodoNotFound) {
				errMessage := fmt.Sprintf("Todo with id %s not found", todoID)
				api.NotFound(w, errMessage)
				return
			}

			api.InternalServerError(w, err.Error())
			return
		}

		api.Success(w, response.FromTodoToResponse(todo))
	}
}
