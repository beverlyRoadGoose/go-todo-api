package item

import (
	"github.com/google/uuid"
	"todo-api/internal/database"
)

type ItemsManager struct{}

var dh = database.NewHandler()

func NewItemsManager() *ItemsManager {
	return &ItemsManager{}
}

func (im *ItemsManager) CreateItem(text string) Item {
	item := Item{
		Id:   uuid.New(),
		Text: text,
	}
	dh.DB().Create(item)
	return item
}
