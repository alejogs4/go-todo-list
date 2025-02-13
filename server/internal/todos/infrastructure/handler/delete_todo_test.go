package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-todo-list.com/m/internal/todos/domain"
)

func TestHandleDeleteTodo(t *testing.T) {
	assertions := require.New(t)

	tests := []struct {
		name           string
		todoRepository domain.Repository
		statusCode     int
	}{
		{
			name: "Delete todo successfully returns 204",
			todoRepository: &domain.RepositoryMock{
				DeleteFunc: func(ctx context.Context, todoID string) error {
					return nil
				},
			},
			statusCode: http.StatusNoContent,
		},
		{
			name: "Delete todo error returns 500",
			todoRepository: &domain.RepositoryMock{
				DeleteFunc: func(ctx context.Context, todoID string) error {
					return errors.New("internal server error")
				},
			},
			statusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := HandleDeleteTodo(tt.todoRepository)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/api/v1/todo/{id}", nil)
			request = mux.SetURLVars(request, map[string]string{"id": "123"})

			handler(recorder, request)
			assertions.Equal(tt.statusCode, recorder.Code)
		})
	}
}
