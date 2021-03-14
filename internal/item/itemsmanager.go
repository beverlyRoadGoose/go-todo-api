package item

import (
	"github.com/google/uuid"
	"todo-api/internal/database"
)

type ItemsManager struct{}

func NewItemsManager() *ItemsManager {
	return &ItemsManager{}
}

func (im *ItemsManager) CreateItem(text string) Item {
	item := Item{
		Id:   uuid.New(),
		Text: text,
	}
	database.DB.Create(item)
	return item
}
