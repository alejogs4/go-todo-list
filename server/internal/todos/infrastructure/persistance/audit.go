package persistance

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type TodoAudit struct {
	ID          int
	TodoID      string
	Action      string
	Description string
	CreatedAt   time.Time
}

const (
	createAction     = "CREATE"
	updateAction     = "UPDATE"
	deleteAction     = "DELETE"
	completeAction   = "COMPLETE"
	unCompleteAction = "UNCOMPLETE"
)

var (
	ErrNotFoundAudit = errors.New("audit not found")
	ErrGettingAudit  = errors.New("error getting audit")
	ErrCreatingAudit = errors.New("error creating audit")
)

//go:generate moq -out todo_audit_mock.go . AuditRepository
type AuditRepository interface {
	Create(ctx context.Context, todo TodoAudit) error
	GetByID(ctx context.Context, id int) (TodoAudit, error)
	GetAll(ctx context.Context) ([]TodoAudit, error)
}

type PostgresAuditRepository struct {
	db *sql.DB
}

func NewPostgresAuditRepository(db *sql.DB) PostgresAuditRepository {
	return PostgresAuditRepository{db: db}
}

func (p PostgresAuditRepository) Create(ctx context.Context, todo TodoAudit) error {
	if _, err := p.db.ExecContext(
		ctx,
		"INSERT INTO todos_audit (todo_id, action, description, created_at) VALUES ($1, $2, $3, $4);",
		todo.TodoID,
		todo.Action,
		todo.Description,
		todo.CreatedAt,
	); err != nil {
		return ErrCreatingAudit
	}

	return nil
}

func (p PostgresAuditRepository) GetByID(ctx context.Context, id int) (TodoAudit, error) {
	var todoAudit TodoAudit

	row := p.db.QueryRowContext(ctx, "SELECT id, todo_id, action, description, created_at FROM todos_audit WHERE id = $1;", id)
	err := row.Scan(&todoAudit.ID, &todoAudit.TodoID, &todoAudit.Action, &todoAudit.Description, &todoAudit.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TodoAudit{}, ErrNotFoundAudit
		}

		return TodoAudit{}, ErrGettingAudit
	}

	return todoAudit, nil
}

func (p PostgresAuditRepository) GetAll(ctx context.Context) ([]TodoAudit, error) {
	var todos []TodoAudit

	rows, err := p.db.QueryContext(ctx, "SELECT id, todo_id, action, description, created_at FROM todos_audit ORDER BY created_at DESC;")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFoundAudit
		}

		return nil, ErrGettingAudit
	}

	for rows.Next() {
		var todoAudit TodoAudit
		err := rows.Scan(&todoAudit.ID, &todoAudit.TodoID, &todoAudit.Action, &todoAudit.Description, &todoAudit.CreatedAt)
		if err != nil {
			return nil, ErrGettingAudit
		}

		todos = append(todos, todoAudit)
	}

	return todos, nil
}
