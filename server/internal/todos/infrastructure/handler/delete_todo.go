package handler

import (
	"net/http"

	"go-todo-list.com/m/internal/shared/api"
	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/use_cases"
)

func HandleDeleteTodo(todoRepository domain.Repository) http.HandlerFunc {
	deleteTodoUseCase := use_cases.NewDeleteTodoUseCase(todoRepository)

	return func(w http.ResponseWriter, r *http.Request) {
		todoID, err := api.GetPathParam(r, "id")
		if err != nil {
			api.InvalidRequest(w, "Missing todo id")
			return
		}

		ctx := r.Context()
		err = deleteTodoUseCase.Execute(ctx, todoID)
		if err != nil {
			api.InternalServerError(w, err.Error())
			return
		}

		api.NoContent(w)
	}
}
