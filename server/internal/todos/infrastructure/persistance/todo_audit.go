package persistance

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-todo-list.com/m/internal/todos/domain"
)

type TodoAuditRepositoryDecorator struct {
	wrappedRepository   domain.Repository
	todoAuditRepository AuditRepository
	logger              *log.Logger
}

func NewTodoAuditRepositoryDecorator(wrappedRepository domain.Repository, todoAuditRepository AuditRepository, logger *log.Logger) TodoAuditRepositoryDecorator {
	return TodoAuditRepositoryDecorator{wrappedRepository: wrappedRepository, todoAuditRepository: todoAuditRepository, logger: logger}
}

func (t TodoAuditRepositoryDecorator) FindBy(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
	return t.wrappedRepository.FindBy(ctx, filters)
}

func (t TodoAuditRepositoryDecorator) Create(ctx context.Context, todo domain.Todo) error {
	err := t.wrappedRepository.Create(ctx, todo)
	if err != nil {
		return err
	}

	audit := TodoAudit{
		TodoID:      todo.ID,
		Action:      "CREATE",
		Description: buildActionDescription(createAction, todo, nil),
		CreatedAt:   time.Now(),
	}
	if err = t.todoAuditRepository.Create(ctx, audit); err != nil {
		t.logger.Printf("Error creating audit record: %v", err)
	}

	return nil
}

func (t TodoAuditRepositoryDecorator) Update(ctx context.Context, todoID string, todo domain.Todo) error {
	oldTodo := t.findTodoByID(ctx, todoID)

	err := t.wrappedRepository.Update(ctx, todoID, todo)
	if err != nil {
		return err
	}

	action := updateAction
	if oldTodo != nil && oldTodo.Completed != todo.Completed {
		if todo.Completed {
			action = completeAction
		} else {
			action = unCompleteAction
		}
	}

	audit := TodoAudit{
		TodoID:      todo.ID,
		Action:      action,
		Description: buildActionDescription(action, todo, oldTodo),
		CreatedAt:   time.Now(),
	}
	if err = t.todoAuditRepository.Create(ctx, audit); err != nil {
		t.logger.Printf("Error creating audit record: %v for update", err)
	}

	return nil
}

func (t TodoAuditRepositoryDecorator) Delete(ctx context.Context, todoID string) error {
	removedTodo, err := t.wrappedRepository.FindBy(ctx, domain.TodoFilters{ID: &todoID})
	if err != nil {
		return err
	}

	err = t.wrappedRepository.Delete(ctx, todoID)
	if err != nil {
		return err
	}

	audit := TodoAudit{
		TodoID:      todoID,
		Action:      deleteAction,
		Description: buildActionDescription(deleteAction, removedTodo[0], nil),
		CreatedAt:   time.Now(),
	}
	if err = t.todoAuditRepository.Create(ctx, audit); err != nil {
		t.logger.Printf("Error creating audit record: %v for delete", err)
	}

	return nil
}

func (t TodoAuditRepositoryDecorator) findTodoByID(ctx context.Context, id string) *domain.Todo {
	todos, err := t.wrappedRepository.FindBy(ctx, domain.TodoFilters{ID: &id})
	if err != nil {
		return nil
	}

	if len(todos) == 0 {
		return nil
	}

	return &todos[0]
}

func buildActionDescription(action string, todo domain.Todo, oldTodo *domain.Todo) string {
	actionMessages := map[string]string{
		"CREATE":     "Todo with title %s created at %s",
		"UPDATE":     "Todo title changed from %s updated at %s to %s",
		"DELETE":     "Todo with title %s deleted at %s",
		"COMPLETE":   "Todo with title %s completed at %s",
		"UNCOMPLETE": "Todo with title %s uncompleted at %s",
	}

	message, ok := actionMessages[action]
	if !ok {
		return fmt.Sprintf("Unknown action %s over todo with title %s", action, todo.Title)
	}

	oldTitle := ""
	if oldTodo != nil {
		oldTitle = oldTodo.Title
	}

	actionArgs := map[string][]any{
		"CREATE":     {todo.Title, time.Now().Format(time.DateTime)},
		"UPDATE":     {oldTitle, time.Now().Format(time.DateTime), todo.Title},
		"DELETE":     {todo.Title, time.Now().Format(time.DateTime)},
		"COMPLETE":   {todo.Title, time.Now().Format(time.DateTime)},
		"UNCOMPLETE": {todo.Title, time.Now().Format(time.DateTime)},
	}

	args, ok := actionArgs[action]
	if !ok {
		return fmt.Sprintf("Unknown action %s over todo with title %s", action, todo.Title)
	}

	return fmt.Sprintf(message, args...)
}
