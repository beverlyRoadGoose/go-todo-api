package item

import (
	"github.com/google/uuid"
	"todo-api/internal/database"
)

type ItemsRepository struct{}

var dh = database.GetHandler()

func NewItemsRepository() *ItemsRepository {
	return &ItemsRepository{}
}

// save saves the given item into the database
func (_ *ItemsRepository) save(i *Item) error {
	if query := dh.DB().Create(&i); query.Error != nil {
		return query.Error
	}
	return nil
}

// update updates the given item in the database
func (_ *ItemsRepository) update(i *Item) error {
	if query := dh.DB().Save(&i); query.Error != nil {
		return query.Error
	}
	return nil
}

// findById returns an item with an ID that matches the given uuid
func (_ *ItemsRepository) findById(id *uuid.UUID) (*Item, error) {
	var item *Item
	if query := dh.DB().First(&item, &id); query.Error != nil {
		return nil, query.Error
	}
	return item, nil
}

// delete deletes an item with an ID that matches the given uuid
func (_ *ItemsRepository) delete(i *Item) error {
	if query := dh.DB().Delete(&i); query.Error != nil {
		return query.Error
	}
	return nil
}

// getAll returns a slice of all items
func (_ *ItemsRepository) getAll() (*[]Item, error) {
	var items *[]Item
	if query := dh.DB().Find(&items); query.Error != nil {
		return nil, query.Error
	}
	return items, nil
}
