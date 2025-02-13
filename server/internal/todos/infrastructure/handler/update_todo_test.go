package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-todo-list.com/m/internal/todos/domain"
)

func TestHandleUpdateTodo(t *testing.T) {
	assertions := require.New(t)

	tests := []struct {
		name               string
		todoRepository     domain.Repository
		requestBody        string
		expectedStatusCode int
	}{
		{
			name: "Update todo successfully returns 204",
			todoRepository: &domain.RepositoryMock{
				UpdateFunc: func(ctx context.Context, todoID string, todo domain.Todo) error {
					return nil
				},
			},
			requestBody:        `{"title": "test", "completed": true}`,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "Update todo with uncomplete successfully returns 204",
			todoRepository: &domain.RepositoryMock{
				UpdateFunc: func(ctx context.Context, todoID string, todo domain.Todo) error {
					return nil
				},
			},
			requestBody:        `{"title": "test", "completed": false}`,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:               "Update todo with empty title returns 400",
			todoRepository:     &domain.RepositoryMock{},
			requestBody:        `{"title": ""}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Update request with wrong body returns 400",
			todoRepository:     &domain.RepositoryMock{},
			requestBody:        `{"wrong": "body"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "Update todo with empty request body returns 400",
			todoRepository:     &domain.RepositoryMock{},
			requestBody:        `{}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Update todo with error returns 500",
			todoRepository: &domain.RepositoryMock{
				UpdateFunc: func(ctx context.Context, todoID string, todo domain.Todo) error {
					return errors.New("internal server error")
				},
			},
			requestBody:        `{"title": "test", "completed": true}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := HandleUpdateTodo(tt.todoRepository)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/api/v1/todo/{id}", strings.NewReader(tt.requestBody))
			request = mux.SetURLVars(request, map[string]string{"id": "123"})

			handler(recorder, request)
			assertions.Equal(tt.expectedStatusCode, recorder.Code)
		})
	}
}
