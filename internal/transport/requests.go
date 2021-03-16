package transport

import (
	"github.com/google/uuid"
	"todo-api/internal/item"
)

type CreateItemRequest struct {
	Text string `json:"text"`
}

type CreateItemResponse struct {
	Item item.Item `json:"item"`
}

type UpdateItemRequest struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
}

type UpdateItemResponse struct {
	Item item.Item `json:"item"`
}

type DeleteItemRequest struct {
	Id uuid.UUID `json:"id"`
}

type DeleteItemResponse struct {
	Deleted bool `json:"deleted"`
}
