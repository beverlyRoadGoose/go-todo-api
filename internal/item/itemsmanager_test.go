package item

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"todo-api/pkg/utils"
)

func TestItemsManager_getItem(t *testing.T) {
	id := uuid.New()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &id).Return(&Item{Id: id}, nil)

	item, err := im.GetItem(id)
	assert.Equal(t, item.Id, id)
	assert.Nil(t, err)
}

func TestItemsManager_GetNonExistingItem(t *testing.T) {
	id := uuid.New()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &id).Return(nil, errors.New("no record found"))

	item, err := im.GetItem(id)
	assert.Nil(t, item)
	assert.Error(t, err)
}

func TestItemsManager_CreateItem(t *testing.T) {
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("save", mock.IsType(&Item{})).Return(nil)
	text := "Test item"

	item, err := im.CreateItem(text)
	assert.True(t, utils.IsValidUUID(item.Id.String()))
	assert.Equal(t, item.Text, text)
	assert.False(t, item.Done)
	assert.Nil(t, err)
}

func TestItemsManager_CreateItemShouldReturnErrorIfErrorHappensDuringDatabaseUpdate(t *testing.T) {
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("save", mock.IsType(&Item{})).Return(errors.New("couldn't write to database"))

	item, err := im.CreateItem("test item")
	assert.Nil(t, item)
	assert.NotNil(t, err)
}

func TestItemsManager_UpdateItem(t *testing.T) {
	item := getTestItem()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &item.Id).Return(item, nil)
	mockRepo.On("update", mock.IsType(&Item{})).Return(nil)

	originalId := item.Id
	updatedText := "New text"
	updatedDone := true

	item, err := im.UpdateItem(item.Id, updatedText, updatedDone)
	assert.Equal(t, item.Id, originalId)
	assert.Equal(t, item.Text, updatedText)
	assert.Equal(t, item.Done, updatedDone)
	assert.Nil(t, err)
}

func TestItemsManager_UpdateNonExistingItem(t *testing.T) {
	id := uuid.New()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &id).Return(nil, errors.New("no record found"))

	item, err := im.UpdateItem(id, "text", true)
	assert.Nil(t, item)
	assert.NotNil(t, err)
}

func TestItemsManager_UpdateItemShouldReturnErrorIfErrorHappensDuringDatabaseUpdate(t *testing.T) {
	item := getTestItem()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &item.Id).Return(item, nil)
	mockRepo.On("update", mock.IsType(&Item{})).Return(errors.New("couldn't write to database"))

	item, err := im.UpdateItem(item.Id, "text", true)
	assert.Nil(t, item)
	assert.NotNil(t, err)
}

func TestItemsManager_DeleteItem(t *testing.T) {
	item := getTestItem()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &item.Id).Return(item, nil)
	mockRepo.On("delete", mock.IsType(&Item{})).Return(nil)

	deleted, err := im.DeleteItem(item.Id)
	assert.True(t, deleted)
	assert.Nil(t, err)
}

func TestItemsManager_DeleteNonExistingItem(t *testing.T) {
	id := uuid.New()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &id).Return(nil, errors.New("no record found"))

	deleted, err := im.DeleteItem(id)
	assert.False(t, deleted)
	assert.NotNil(t, err)
}

func TestItemsManager_DeleteShouldReturnErrorIfErrorHappensDuringDatabaseUpdate(t *testing.T) {
	item := getTestItem()
	mockRepo := &MockItemsRepository{}
	im := NewItemsManager(mockRepo)
	mockRepo.On("findById", &item.Id).Return(item, nil)
	mockRepo.On("delete", mock.IsType(&Item{})).Return(errors.New("couldn't write to database"))

	deleted, err := im.DeleteItem(item.Id)
	assert.False(t, deleted)
	assert.NotNil(t, err)
}

func getTestItem() *Item {
	return &Item{
		Id:   uuid.New(),
		Text: "Test item",
		Done: false,
	}
}
