package internal

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-api/internal/item"
)

func TestTodoApiService_GetItem(t *testing.T) {
	id := uuid.New()
	mockRepo := &item.MockItemsRepository{}
	mockRepo.On("findById", &id).Return(&item.Item{Id: id}, nil)
	im := item.NewItemsManager(mockRepo)
	s := NewTodoApiService(im)

	it, err := s.GetItem(id)
	assert.Equal(t, it.Id, id)
	assert.Nil(t, err)
}

func TestTodoApiService_GetNonExistingItem(t *testing.T) {
	id := uuid.New()
	mockRepo := &item.MockItemsRepository{}
	mockRepo.On("findById", &id).Return(nil, errors.New("no record found"))
	im := item.NewItemsManager(mockRepo)
	s := NewTodoApiService(im)

	it, err := s.GetItem(id)
	assert.Nil(t, it)
	assert.NotNil(t, err)
}

func getTestItem() *item.Item {
	return &item.Item{
		Id:   uuid.New(),
		Text: "Test item",
		Done: false,
	}
}
