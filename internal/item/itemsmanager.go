package item

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type ItemsManager struct {
	repository Repository
}

func NewItemsManager(r Repository) *ItemsManager {
	return &ItemsManager{repository: r}
}

func (im *ItemsManager) GetItem(id uuid.UUID) (*Item, error) {
	item, err := im.repository.findById(&id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (im *ItemsManager) CreateItem(text string) (*Item, error) {
	item := &Item{Id: uuid.New(), Text: text}
	if err := im.repository.save(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (im *ItemsManager) UpdateItem(id uuid.UUID, text string, done bool) (*Item, error) {
	item, err := im.GetItem(id)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"id": id, "text": text, "done": done}).Info("Updating item")
	item.Text = text
	item.Done = done
	if err := im.repository.update(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (im *ItemsManager) DeleteItem(id uuid.UUID) (bool, error) {
	item, err := im.GetItem(id)
	if err != nil {
		return false, err
	}
	log.WithFields(log.Fields{"id": id}).Info("Deleting item")
	if err := im.repository.delete(item); err != nil {
		return false, err
	}
	return true, nil
}
