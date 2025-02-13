package use_cases

import (
	"context"

	"go-todo-list.com/m/internal/todos/domain"
)

type CreateTodoUseCase struct {
	repository  domain.Repository
	idGenerator domain.IDGenerator
}

func NewCreateTodosUseCase(repository domain.Repository, idGenerator domain.IDGenerator) CreateTodoUseCase {
	return CreateTodoUseCase{repository: repository, idGenerator: idGenerator}
}

func (u CreateTodoUseCase) Execute(ctx context.Context, title string) (string, error) {
	uuid := u.idGenerator.Generate()

	todo, err := domain.NewTodo(uuid, title)
	if err != nil {
		return "", err
	}

	err = u.repository.Create(ctx, todo)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
