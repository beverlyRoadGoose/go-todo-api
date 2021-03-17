package item

import (
	"github.com/google/uuid"
	"todo-api/internal/database"
)

type Repository interface {
	save(i *Item) error
	update(i *Item) error
	findById(id *uuid.UUID) (*Item, error)
	delete(i *Item) error
	getAll() (*[]Item, error)
}

type ItemsRepository struct {
	databaseHandler *database.Handler
}

func NewItemsRepository() *ItemsRepository {
	return &ItemsRepository{
		databaseHandler: database.GetHandler(),
	}
}

// save saves the given item into the database
func (r *ItemsRepository) save(i *Item) error {
	if query := r.databaseHandler.DB().Create(&i); query.Error != nil {
		return query.Error
	}
	return nil
}

// update updates the given item in the database
func (r *ItemsRepository) update(i *Item) error {
	if query := r.databaseHandler.DB().Save(&i); query.Error != nil {
		return query.Error
	}
	return nil
}

// findById returns an item with an ID that matches the given uuid
func (r *ItemsRepository) findById(id *uuid.UUID) (*Item, error) {
	var item *Item
	if query := r.databaseHandler.DB().First(&item, &id); query.Error != nil {
		return nil, query.Error
	}
	return item, nil
}

// delete deletes an item with an ID that matches the given uuid
func (r *ItemsRepository) delete(i *Item) error {
	if query := r.databaseHandler.DB().Delete(&i); query.Error != nil {
		return query.Error
	}
	return nil
}

// getAll returns a slice of all items
func (r *ItemsRepository) getAll() (*[]Item, error) {
	var items *[]Item
	if query := r.databaseHandler.DB().Find(&items); query.Error != nil {
		return nil, query.Error
	}
	return items, nil
}
