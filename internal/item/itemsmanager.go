package item

import (
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"todo-api/internal/database"
)

type ItemsManager struct{}

var dh = database.NewHandler()

func NewItemsManager() *ItemsManager {
	return &ItemsManager{}
}

func itemExists(id uuid.UUID) bool {
	if query := dh.DB().First(&Item{}, id); query.Error != nil {
		return false
	}
	return true
}

func (im *ItemsManager) CreateItem(text string) *Item {
	item := Item{
		Id:   uuid.New(),
		Text: text,
	}
	dh.DB().Create(item)
	return &item
}

func (im *ItemsManager) UpdateItem(id uuid.UUID, text string, done bool) (*Item, error) {
	if !itemExists(id) {
		return nil, errors.New("no item found with the given id")
	}
	log.WithFields(log.Fields{"id": id, "text": text, "done": done}).Info("Updating item")
	item := &Item{}
	dh.DB().First(&item, id)
	item.Text = text
	item.Done = done
	dh.DB().Save(&item)
	return item, nil
}
