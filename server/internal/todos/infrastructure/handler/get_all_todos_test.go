package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go-todo-list.com/m/internal/shared/test_utils"
	"go-todo-list.com/m/internal/shared/utils"
	"go-todo-list.com/m/internal/todos/domain"
	"go-todo-list.com/m/internal/todos/infrastructure/persistance"
)

func TestHandleGetAllTodos(t *testing.T) {
	assertions := require.New(t)

	tests := []struct {
		name               string
		todoRepository     domain.Repository
		completed          *string
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name: "Get all todos successfully",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return []domain.Todo{
						{
							ID:          "123",
							Title:       "test",
							Completed:   false,
							CompletedAt: nil,
							CreatedAt:   time.Time{},
							UpdatedAt:   time.Time{},
						},
						{
							ID:          "456",
							Title:       "test2",
							Completed:   true,
							CompletedAt: &time.Time{},
							CreatedAt:   time.Time{},
							UpdatedAt:   time.Time{},
						},
					}, nil
				},
			},
			completed:          nil,
			expectedResponse:   `{"content":[{"id":"123","title":"test","completed":false,"created_at":"0001-01-01 00:00:00","completed_at":null,"updated_at":"0001-01-01 00:00:00"},{"id":"456","title":"test2","completed":true,"created_at":"0001-01-01 00:00:00","completed_at":"0001-01-01 00:00:00","updated_at":"0001-01-01 00:00:00"}]}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Get all completed todos successfully with query param",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return []domain.Todo{
						{
							ID:          "456",
							Title:       "test2",
							Completed:   true,
							CompletedAt: &time.Time{},
							CreatedAt:   time.Time{},
							UpdatedAt:   time.Time{},
						},
					}, nil
				},
			},
			completed:          utils.Ptr("true"),
			expectedResponse:   `{"content":[{"id":"456","title":"test2","completed":true,"created_at":"0001-01-01 00:00:00","completed_at":"0001-01-01 00:00:00","updated_at":"0001-01-01 00:00:00"}]}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Get all not completed todos successfully with query param false",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return []domain.Todo{
						{
							ID:          "123",
							Title:       "test",
							Completed:   false,
							CompletedAt: nil,
							CreatedAt:   time.Time{},
							UpdatedAt:   time.Time{},
						},
					}, nil
				},
			},
			completed:          utils.Ptr("false"),
			expectedResponse:   `{"content":[{"id":"123","title":"test","completed":false,"created_at":"0001-01-01 00:00:00","completed_at":null,"updated_at":"0001-01-01 00:00:00"}]}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "error when query param is invalid",
			todoRepository:     &domain.RepositoryMock{},
			completed:          utils.Ptr("invalid"),
			expectedResponse:   `{"app_code":"INVALID_REQUEST", "message":"Invalid query param it must be true or false"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "error when repository returns error return 500",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return nil, errors.New("error finding todos")
				},
			},
			completed:          nil,
			expectedResponse:   `{"app_code":"INTERNAL_SERVER_ERROR", "message":"error finding todos"}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "error when repository returns not found and completed query param is not passed return 200",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return nil, domain.ErrTodoNotFound
				},
			},
			completed:          nil,
			expectedResponse:   `{"content":[]}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "error when repository returns not found and completed query param is passed return 404",
			todoRepository: &domain.RepositoryMock{
				FindByFunc: func(ctx context.Context, filters domain.TodoFilters) ([]domain.Todo, error) {
					return nil, domain.ErrTodoNotFound
				},
			},
			completed:          utils.Ptr("true"),
			expectedResponse:   `{"app_code":"NOT_FOUND", "message":"Todos not found"}`,
			expectedStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := HandleGetAllTodos(tt.todoRepository)
			request, err := http.NewRequest(http.MethodGet, "/api/v1/todo", nil)
			assertions.NoError(err)
			q := request.URL.Query()
			if tt.completed != nil {
				q.Add("completed", *tt.completed)
			}
			request.URL.RawQuery = q.Encode()

			recorder := httptest.NewRecorder()

			handler.ServeHTTP(recorder, request)

			assertions.Equal(tt.expectedStatusCode, recorder.Code)
			assertions.JSONEq(tt.expectedResponse, recorder.Body.String())
		})
	}
}

func TestIntegrationHandleGetAllTodos(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("skipping integration test")
	}

	assertions := require.New(t)

	createdAt := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	testCases := []struct {
		name               string
		completed          *string
		preconditions      func(ctx context.Context, repository domain.Repository)
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name:      "Get all todos successfully",
			completed: nil,
			preconditions: func(ctx context.Context, repository domain.Repository) {
				err := repository.Create(context.Background(), domain.Todo{
					ID:        "1",
					Title:     "test",
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				})
				assertions.NoError(err)

				err = repository.Create(context.Background(), domain.Todo{
					ID:        "2",
					Title:     "test2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
				})
				assertions.NoError(err)
			},
			expectedResponse:   `{"content":[{"id":"2","title":"test2","completed":false,"created_at":"2021-02-01 00:00:00","completed_at":null,"updated_at":"2021-02-01 00:00:00"},{"id":"1","title":"test","completed":false,"created_at":"2021-01-01 00:00:00","completed_at":null,"updated_at":"2021-01-01 00:00:00"}]}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:      "Get all completed todos successfully with query param",
			completed: utils.Ptr("true"),
			preconditions: func(ctx context.Context, repository domain.Repository) {
				err := repository.Create(ctx, domain.Todo{
					ID:          "3",
					Title:       "test2",
					Completed:   true,
					CreatedAt:   time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					CompletedAt: utils.Ptr(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
				})
				assertions.NoError(err)

				err = repository.Create(ctx, domain.Todo{
					ID:        "4",
					Title:     "test",
					CreatedAt: createdAt,
					UpdatedAt: createdAt,
				})
				assertions.NoError(err)
			},
			expectedResponse:   `{"content":[{"id":"3","title":"test2","completed":true,"created_at":"2021-02-01 00:00:00","completed_at":"2021-02-01 00:00:00","updated_at":"2021-02-01 00:00:00"}]}`,
			expectedStatusCode: http.StatusOK,
		},
	}

	ctx := context.Background()
	database, err := test_utils.GenerateTestPostgresDB(ctx, "init.sql", "todos")
	assertions.NoError(err)

	t.Cleanup(func() {
		assertions.NoError(database.Container.Terminate(ctx))
	})

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repository := persistance.NewTodosPostgresRepository(database.DB)
			tt.preconditions(ctx, repository)

			handler := HandleGetAllTodos(repository)
			request, err := http.NewRequest(http.MethodGet, "/api/v1/todo", nil)
			assertions.NoError(err)
			q := request.URL.Query()
			if tt.completed != nil {
				q.Add("completed", *tt.completed)
			}
			request.URL.RawQuery = q.Encode()

			recorder := httptest.NewRecorder()

			handler(recorder, request)

			assertions.Equal(tt.expectedStatusCode, recorder.Code)
			assertions.JSONEq(tt.expectedResponse, recorder.Body.String())
		})
	}
}
