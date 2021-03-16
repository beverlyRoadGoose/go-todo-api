package internal

import (
	"github.com/google/uuid"
	"todo-api/internal/item"
)

type Service interface {
	CreateItem(text string) *item.Item
	UpdateItem(id uuid.UUID, text string, done bool) (*item.Item, error)
	DeleteItem(id uuid.UUID) (bool, error)
}

type todoApiService struct{}

var itemsManager = item.NewItemsManager()

func NewTodoApiService() Service {
	return &todoApiService{}
}

func (t *todoApiService) CreateItem(text string) *item.Item {
	return itemsManager.CreateItem(text)
}

func (t *todoApiService) UpdateItem(id uuid.UUID, text string, done bool) (*item.Item, error) {
	return itemsManager.UpdateItem(id, text, done)
}

func (t *todoApiService) DeleteItem(id uuid.UUID) (bool, error) {
	return itemsManager.DeleteItem(id)
}
