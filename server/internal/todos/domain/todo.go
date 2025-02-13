package domain

import "time"

type Todo struct {
	ID          string
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	UpdatedAt   time.Time
}

func NewTodo(id string, title string) (Todo, error) {
	if title == "" {
		return Todo{}, ErrEmptyTodoTitle
	}

	if id == "" {
		return Todo{}, ErrNotProvidedID
	}

	return Todo{
		ID:          id,
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CompletedAt: nil,
	}, nil
}

func (t *Todo) Complete() {
	t.Completed = true

	now := time.Now()
	t.CompletedAt = &now
	t.UpdatedAt = now
}

func (t *Todo) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
	t.UpdatedAt = time.Now()
}
