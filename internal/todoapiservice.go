package internal

import (
	"github.com/google/uuid"
	"todo-api/internal/item"
)

type Service interface {
	CreateItem(text string) item.Item
	UpdateItem(text string, done bool) item.Item
	DeleteItem(id uuid.UUID)
}

type todoApiService struct{}

var itemsManager = item.NewItemsManager()

func NewTodoApiService() Service {
	return &todoApiService{}
}

func (t *todoApiService) CreateItem(text string) item.Item {
	return itemsManager.CreateItem(text)
}

func (t *todoApiService) UpdateItem(text string, done bool) item.Item {
	panic("implement me")
}

func (t *todoApiService) DeleteItem(id uuid.UUID) {
	panic("implement me")
}
