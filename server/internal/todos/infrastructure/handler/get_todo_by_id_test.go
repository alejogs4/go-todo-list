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

func TestHandleGetTodoByID(t *testing.T) {
	assertions := require.New(t)

	tests := []struct {
		name               string
		todoRepository     domain.Repository
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name: "Get todo by id successfully returns 200",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return []domain.Todo{
						{
							ID:        "123",
							Title:     "test",
							Completed: false,
						},
					}, nil
				},
			},
			expectedResponse:   `{"content":{"id":"123","title":"test","completed":false,"created_at":"0001-01-01 00:00:00","completed_at":null,"updated_at":"0001-01-01 00:00:00"}}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Get todo by id not found returns 404",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return nil, domain.ErrTodoNotFound
				},
			},
			expectedResponse:   `{"app_code":"NOT_FOUND", "message":"Todo with id 123 not found"}`,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "Get todo by id internal server error returns 500",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return nil, errors.New("internal server error")
				},
			},
			expectedResponse:   `{"app_code":"INTERNAL_SERVER_ERROR", "message":"internal server error"}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := HandleGetTodoByID(tt.todoRepository)
			path := "/api/v1/todo/{id}"
			req, err := http.NewRequest(http.MethodGet, path, nil)
			req = mux.SetURLVars(req, map[string]string{"id": "123"})
			assertions.NoError(err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			assertions.Equal(tt.expectedStatusCode, rr.Code)
			assertions.JSONEq(tt.expectedResponse, rr.Body.String())
		})
	}
}
