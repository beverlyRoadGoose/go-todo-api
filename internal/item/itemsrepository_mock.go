package item

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockItemsRepository struct {
	mock.Mock
}

func (m *MockItemsRepository) save(i *Item) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *MockItemsRepository) update(i *Item) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *MockItemsRepository) findById(id *uuid.UUID) (*Item, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Item), args.Error(1)
}

func (m *MockItemsRepository) delete(i *Item) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *MockItemsRepository) getAll() (*[]Item, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]Item), args.Error(1)
}
