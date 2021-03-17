package internal

import (
	"github.com/google/uuid"
	"todo-api/internal/item"
)

type Service interface {
	GetItem(id uuid.UUID) (*item.Item, error)
	CreateItem(text string) (*item.Item, error)
	UpdateItem(id uuid.UUID, text string, done bool) (*item.Item, error)
	DeleteItem(id uuid.UUID) (bool, error)
}

type todoApiService struct{}

var itemsManager = item.NewItemsManager()

func NewTodoApiService() Service {
	return &todoApiService{}
}

func (_ *todoApiService) GetItem(id uuid.UUID) (*item.Item, error) {
	return itemsManager.GetItem(id)
}

func (_ *todoApiService) CreateItem(text string) (*item.Item, error) {
	return itemsManager.CreateItem(text)
}

func (_ *todoApiService) UpdateItem(id uuid.UUID, text string, done bool) (*item.Item, error) {
	return itemsManager.UpdateItem(id, text, done)
}

func (_ *todoApiService) DeleteItem(id uuid.UUID) (bool, error) {
	return itemsManager.DeleteItem(id)
}
