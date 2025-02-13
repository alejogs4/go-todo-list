package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go-todo-list.com/m/internal/todos/domain"
)

func TestHandleCreateTodo(t *testing.T) {
	assertions := require.New(t)

	tests := []struct {
		name               string
		todoRepository     domain.Repository
		idGenerator        domain.IDGenerator
		todoRequest        string
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name: "Create todo successfully",
			todoRepository: &domain.RepositoryMock{
				CreateFunc: func(ctx context.Context, todo domain.Todo) error {
					return nil
				},
			},
			idGenerator: &domain.IDGeneratorMock{
				GenerateFunc: func() string {
					return "123"
				},
			},
			todoRequest:        `{"title": "test"}`,
			expectedResponse:   `{"content": {"id": "123"}}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Create todo with empty title returns 400",
			todoRepository: &domain.RepositoryMock{
				CreateFunc: func(ctx context.Context, todo domain.Todo) error {
					return domain.ErrEmptyTodoTitle
				},
			},
			idGenerator: &domain.IDGeneratorMock{
				GenerateFunc: func() string {
					return "123"
				},
			},
			todoRequest:        `{"title": ""}`,
			expectedResponse:   `{"app_code":"INVALID_REQUEST", "message":"todo title cannot be empty"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create todo with empty id returns 400",
			todoRepository: &domain.RepositoryMock{
				CreateFunc: func(ctx context.Context, todo domain.Todo) error {
					return domain.ErrNotProvidedID
				},
			},
			idGenerator: &domain.IDGeneratorMock{
				GenerateFunc: func() string {
					return ""
				},
			},
			todoRequest:        `{"title": "test"}`,
			expectedResponse:   `{"app_code":"INVALID_REQUEST", "message":"id not provided"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := HandleCreateTodo(tt.todoRepository, tt.idGenerator)
			responseRecorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/api/v1/todo", strings.NewReader(tt.todoRequest))

			handler(responseRecorder, request)

			assertions.Equal(tt.expectedStatusCode, responseRecorder.Code)
			assertions.JSONEq(tt.expectedResponse, responseRecorder.Body.String())
		})
	}
}
