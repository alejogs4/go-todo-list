package use_cases

import (
	"context"

	"go-todo-list.com/m/internal/todos/domain"
)

type UpdateTodoUseCase struct {
	repository domain.Repository
}

func NewUpdateTodoUseCase(repository domain.Repository) UpdateTodoUseCase {
	return UpdateTodoUseCase{repository: repository}
}

func (u UpdateTodoUseCase) Execute(ctx context.Context, todoID string, title string, completed bool) error {
	todo, err := domain.NewTodo(todoID, title)
	if err != nil {
		return err
	}

	if completed {
		todo.Complete()
	} else {
		todo.Uncomplete()
	}

	return u.repository.Update(ctx, todoID, todo)
}
