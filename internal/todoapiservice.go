package internal

import (
	"github.com/google/uuid"
	"todo-api/internal/todo"
)

type Service interface {
	CreateItem(text string) todo.Item
	UpdateItem(text string, done bool) todo.Item
	DeleteItem(id uuid.UUID)
}

type todoApiService struct{}

func NewTodoApiService() Service {
	return &todoApiService{}
}

func (t todoApiService) CreateItem(text string) todo.Item {
	panic("implement me")
}

func (t todoApiService) UpdateItem(text string, done bool) todo.Item {
	panic("implement me")
}

func (t todoApiService) DeleteItem(id uuid.UUID) {
	panic("implement me")
}
