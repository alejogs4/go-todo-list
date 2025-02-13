package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"go-todo-list.com/m/internal/shared/api"
	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/request"
	"go-todo-list.com/m/internal/todos/use_cases"
)

func HandleUpdateTodo(todoRepository domain.Repository) http.HandlerFunc {
	updateTodo := use_cases.NewUpdateTodoUseCase(todoRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		todoID, err := api.GetPathParam(r, "id")
		if err != nil {
			api.InvalidRequest(w, "Missing todo id")
			return
		}

		requestTodo, err := updatedTodoFromRequest(r)
		if err != nil {
			api.InvalidRequest(w, "Invalid request body")
			return
		}

		ctx := r.Context()
		err = updateTodo.Execute(ctx, todoID, requestTodo.Title, requestTodo.Completed)
		if err != nil {
			if errors.Is(err, domain.ErrEmptyTodoTitle) || errors.Is(err, domain.ErrNotProvidedID) {
				api.InvalidRequest(w, err.Error())
				return
			}

			api.InternalServerError(w, err.Error())
		}

		api.NoContent(w)
	}
}

func updatedTodoFromRequest(r *http.Request) (request.UpdateTodoRequest, error) {
	var requestBody request.UpdateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return request.UpdateTodoRequest{}, err
	}

	return requestBody, nil
}
