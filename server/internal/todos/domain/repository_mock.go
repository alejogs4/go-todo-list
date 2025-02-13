// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package domain

import (
	"context"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			CreateFunc: func(ctx context.Context, todo Todo) error {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx context.Context, todoID string) error {
//				panic("mock out the Delete method")
//			},
//			FindByFunc: func(ctx context.Context, filters TodoFilters) ([]Todo, error) {
//				panic("mock out the FindBy method")
//			},
//			UpdateFunc: func(ctx context.Context, todoID string, todo Todo) error {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, todo Todo) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, todoID string) error

	// FindByFunc mocks the FindBy method.
	FindByFunc func(ctx context.Context, filters TodoFilters) ([]Todo, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, todoID string, todo Todo) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Todo is the todo argument value.
			Todo Todo
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TodoID is the todoID argument value.
			TodoID string
		}
		// FindBy holds details about calls to the FindBy method.
		FindBy []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filters is the filters argument value.
			Filters TodoFilters
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TodoID is the todoID argument value.
			TodoID string
			// Todo is the todo argument value.
			Todo Todo
		}
	}
	lockCreate sync.RWMutex
	lockDelete sync.RWMutex
	lockFindBy sync.RWMutex
	lockUpdate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *RepositoryMock) Create(ctx context.Context, todo Todo) error {
	if mock.CreateFunc == nil {
		panic("RepositoryMock.CreateFunc: method is nil but Repository.Create was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Todo Todo
	}{
		Ctx:  ctx,
		Todo: todo,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, todo)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedRepository.CreateCalls())
func (mock *RepositoryMock) CreateCalls() []struct {
	Ctx  context.Context
	Todo Todo
} {
	var calls []struct {
		Ctx  context.Context
		Todo Todo
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *RepositoryMock) Delete(ctx context.Context, todoID string) error {
	if mock.DeleteFunc == nil {
		panic("RepositoryMock.DeleteFunc: method is nil but Repository.Delete was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		TodoID string
	}{
		Ctx:    ctx,
		TodoID: todoID,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, todoID)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedRepository.DeleteCalls())
func (mock *RepositoryMock) DeleteCalls() []struct {
	Ctx    context.Context
	TodoID string
} {
	var calls []struct {
		Ctx    context.Context
		TodoID string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// FindBy calls FindByFunc.
func (mock *RepositoryMock) FindBy(ctx context.Context, filters TodoFilters) ([]Todo, error) {
	if mock.FindByFunc == nil {
		panic("RepositoryMock.FindByFunc: method is nil but Repository.FindBy was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Filters TodoFilters
	}{
		Ctx:     ctx,
		Filters: filters,
	}
	mock.lockFindBy.Lock()
	mock.calls.FindBy = append(mock.calls.FindBy, callInfo)
	mock.lockFindBy.Unlock()
	return mock.FindByFunc(ctx, filters)
}

// FindByCalls gets all the calls that were made to FindBy.
// Check the length with:
//
//	len(mockedRepository.FindByCalls())
func (mock *RepositoryMock) FindByCalls() []struct {
	Ctx     context.Context
	Filters TodoFilters
} {
	var calls []struct {
		Ctx     context.Context
		Filters TodoFilters
	}
	mock.lockFindBy.RLock()
	calls = mock.calls.FindBy
	mock.lockFindBy.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *RepositoryMock) Update(ctx context.Context, todoID string, todo Todo) error {
	if mock.UpdateFunc == nil {
		panic("RepositoryMock.UpdateFunc: method is nil but Repository.Update was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		TodoID string
		Todo   Todo
	}{
		Ctx:    ctx,
		TodoID: todoID,
		Todo:   todo,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, todoID, todo)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedRepository.UpdateCalls())
func (mock *RepositoryMock) UpdateCalls() []struct {
	Ctx    context.Context
	TodoID string
	Todo   Todo
} {
	var calls []struct {
		Ctx    context.Context
		TodoID string
		Todo   Todo
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
