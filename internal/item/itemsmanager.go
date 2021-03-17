package item

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type ItemsManager struct{}

var repository = NewItemsRepository()

func NewItemsManager() *ItemsManager {
	return &ItemsManager{}
}

func getItem(id uuid.UUID) (*Item, error) {
	item, err := repository.findById(&id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (_ *ItemsManager) GetItem(id uuid.UUID) (*Item, error) {
	return getItem(id)
}

func (_ *ItemsManager) CreateItem(text string) (*Item, error) {
	item := &Item{Id: uuid.New(), Text: text}
	if err := repository.save(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (_ *ItemsManager) UpdateItem(id uuid.UUID, text string, done bool) (*Item, error) {
	item, err := getItem(id)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"id": id, "text": text, "done": done}).Info("Updating item")
	item.Text = text
	item.Done = done
	if err := repository.update(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (_ *ItemsManager) DeleteItem(id uuid.UUID) (bool, error) {
	item, err := getItem(id)
	if err != nil {
		return false, err
	}
	log.WithFields(log.Fields{"id": id}).Info("Deleting item")
	if err := repository.delete(item); err != nil {
		return false, err
	}
	return true, nil
}
