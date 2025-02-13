package persistance

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"go-todo-list.com/m/internal/todos/domain"
)

type TodosPostgresRepository struct {
	db *sql.DB
}

func NewTodosPostgresRepository(db *sql.DB) TodosPostgresRepository {
	return TodosPostgresRepository{db: db}
}

var (
	errTodoDatabase = errors.New("error with database")
	errCreatingTodo = errors.New("error creating todo")
	errUpdatingTodo = errors.New("error updating todo")
	errDeletingTodo = errors.New("error deleting todo")
	errParsingTodo  = errors.New("error parsing todo")
)

func (t TodosPostgresRepository) FindBy(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
	query, args := buildFindQuery(filters)

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Todo{}, domain.ErrTodoNotFound
		}

		return nil, errTodoDatabase
	}

	defer rows.Close()
	todos := make([]domain.Todo, 0)

	for rows.Next() {
		var todo domain.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, errParsingTodo
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t TodosPostgresRepository) Create(ctx context.Context, todo domain.Todo) error {
	if _, err := t.db.ExecContext(
		ctx,
		"INSERT INTO todos (id, title, completed, created_at, updated_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6);",
		todo.ID,
		todo.Title,
		todo.Completed,
		todo.CreatedAt,
		todo.UpdatedAt,
		todo.CompletedAt,
	); err != nil {
		return errCreatingTodo
	}

	return nil
}

func (t TodosPostgresRepository) Update(ctx context.Context, todoID string, todo domain.Todo) error {
	if _, err := t.db.ExecContext(
		ctx,
		"UPDATE todos SET title = $1, completed = $2, updated_at = $3, completed_at = $4 WHERE id = $5;",
		todo.Title,
		todo.Completed,
		todo.UpdatedAt,
		todo.CompletedAt,
		todoID,
	); err != nil {
		return errUpdatingTodo
	}

	return nil
}

func (t TodosPostgresRepository) Delete(ctx context.Context, todoID string) error {
	if _, err := t.db.ExecContext(ctx, "DELETE FROM todos WHERE id = $1;", todoID); err != nil {
		return errDeletingTodo
	}

	return nil
}

func buildFindQuery(filters domain.TodoFilters) (string, []interface{}) {
	args := make([]interface{}, 0)
	buf := strings.Builder{}
	buf.WriteString("SELECT id, title, completed, created_at, updated_at, completed_at FROM todos")
	if filters.ID != nil || filters.Completed != nil {
		buf.WriteString(" WHERE ")
	}

	if filters.ID != nil {
		buf.WriteString("id = $1")
		args = append(args, *filters.ID)
	}

	if filters.Completed != nil && filters.ID != nil {
		buf.WriteString(" AND ")
	}

	if filters.Completed != nil {
		completedPosition := len(args) + 1
		buf.WriteString(fmt.Sprintf("completed = $%d", completedPosition))
		args = append(args, *filters.Completed)
	}

	buf.WriteString(" ORDER BY created_at DESC")
	buf.WriteString(";")

	return buf.String(), args
}
