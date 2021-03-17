package item

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemsManager_getItem(t *testing.T) {
	testId := uuid.New()
	mockRepo := new(MockItemsRepository)
	mockRepo.On("findById", &testId).Return(&Item{Id: testId}, nil)

	itemsManager := NewItemsManager(mockRepo)
	item, err := itemsManager.GetItem(testId)

	assert.Equal(t, item.Id, testId, "Returned item id should match test id")
	assert.Nil(t, err, "err should be nil when match is found successfully")
}
