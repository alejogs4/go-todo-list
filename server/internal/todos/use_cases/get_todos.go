package use_cases

import (
	"context"

	"go-todo-list.com/m/internal/todos/domain"
)

type GetTodosUseCase struct {
	repository domain.Repository
}

func NewGetTodosUseCase(repository domain.Repository) GetTodosUseCase {
	return GetTodosUseCase{repository: repository}
}

func (u GetTodosUseCase) Execute(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
	return u.repository.FindBy(ctx, filters)
}
