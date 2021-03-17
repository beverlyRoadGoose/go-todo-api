package item

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
	log.Info("Invoking mocked items repository findById method")
	args := m.Called(id)
	return args.Get(0).(*Item), args.Error(1)
}

func (m *MockItemsRepository) delete(i *Item) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *MockItemsRepository) getAll() (*[]Item, error) {
	args := m.Called()
	return args.Get(0).(*[]Item), args.Error(1)
}
