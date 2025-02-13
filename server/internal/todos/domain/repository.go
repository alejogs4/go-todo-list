package domain

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	FindBy(ctx context.Context, filters TodoFilters) ([]Todo, error)
	Create(ctx context.Context, todo Todo) error
	Update(ctx context.Context, todoID string, todo Todo) error
	Delete(ctx context.Context, todoID string) error
}

type TodoFilters struct {
	Completed *bool
	ID        *string
}
