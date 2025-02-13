package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"go-todo-list.com/m/internal/shared/api"
	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/request"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/response"
	"go-todo-list.com/m/internal/todos/use_cases"
)

func HandleCreateTodo(todoRepository domain.Repository, idGenerator domain.IDGenerator) http.HandlerFunc {
	createTodoUseCase := use_cases.NewCreateTodosUseCase(todoRepository, idGenerator)

	return func(w http.ResponseWriter, r *http.Request) {
		requestTodo, err := newTodoFromRequest(r)
		if err != nil {
			api.InvalidRequest(w, "Invalid request body")
			return
		}

		ctx := r.Context()
		uuid, err := createTodoUseCase.Execute(ctx, requestTodo.Title)
		if err != nil {
			if errors.Is(err, domain.ErrEmptyTodoTitle) || errors.Is(err, domain.ErrNotProvidedID) {
				api.InvalidRequest(w, err.Error())
				return
			}

			api.InternalServerError(w, err.Error())
			return
		}

		api.Created(w, response.CreateTodoResponse{ID: uuid})
	}
}

func newTodoFromRequest(r *http.Request) (request.CreateTodoRequest, error) {
	var requestBody request.CreateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return request.CreateTodoRequest{}, err
	}

	return requestBody, nil
}
