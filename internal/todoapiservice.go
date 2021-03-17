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

type todoApiService struct {
	itemsManager *item.ItemsManager
}

func NewTodoApiService() Service {
	return &todoApiService{
		itemsManager: item.NewItemsManager(item.NewItemsRepository()),
	}
}

func (s *todoApiService) GetItem(id uuid.UUID) (*item.Item, error) {
	return s.itemsManager.GetItem(id)
}

func (s *todoApiService) CreateItem(text string) (*item.Item, error) {
	return s.itemsManager.CreateItem(text)
}

func (s *todoApiService) UpdateItem(id uuid.UUID, text string, done bool) (*item.Item, error) {
	return s.itemsManager.UpdateItem(id, text, done)
}

func (s *todoApiService) DeleteItem(id uuid.UUID) (bool, error) {
	return s.itemsManager.DeleteItem(id)
}
