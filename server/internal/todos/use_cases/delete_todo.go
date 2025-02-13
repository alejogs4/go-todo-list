package use_cases

import (
	"context"

	"go-todo-list.com/m/internal/todos/domain"
)

type DeleteTodoUseCase struct {
	repository domain.Repository
}

func NewDeleteTodoUseCase(repository domain.Repository) DeleteTodoUseCase {
	return DeleteTodoUseCase{repository: repository}
}

func (u DeleteTodoUseCase) Execute(ctx context.Context, todoID string) error {
	return u.repository.Delete(ctx, todoID)
}
