package domain

import "errors"

var (
	ErrEmptyTodoTitle = errors.New("todo title cannot be empty")
	ErrNotProvidedID  = errors.New("id not provided")
	ErrTodoNotFound   = errors.New("todo not found")
)
