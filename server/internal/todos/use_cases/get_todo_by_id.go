package use_cases

import (
	"context"

	"go-todo-list.com/m/internal/todos/domain"
)

type GetTodosByIDUseCase struct {
	repository domain.Repository
}

func NewGetTodosByIDUseCase(repository domain.Repository) GetTodosByIDUseCase {
	return GetTodosByIDUseCase{repository: repository}
}

func (u GetTodosByIDUseCase) Execute(ctx context.Context, filters domain.TodoFilters) (domain.Todo, error) {
	todos, err := u.repository.FindBy(ctx, filters)
	if err != nil {
		return domain.Todo{}, err
	}

	if len(todos) == 0 {
		return domain.Todo{}, domain.ErrTodoNotFound
	}

	return todos[0], nil
}
