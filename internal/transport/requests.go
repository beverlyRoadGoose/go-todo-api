package transport

import (
	"github.com/google/uuid"
	"todo-api/internal/todo"
)

type CreateItemRequest struct {
	Text string `json:"text"`
}

type CreateItemResponse struct {
	Item todo.Item `json:"item"`
}

type UpdateItemRequest struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type UpdateItemResponse struct {
	Item todo.Item `json:"item"`
}

type DeleteItemRequest struct {
	Id uuid.UUID `json:"id"`
}

type DeleteItemResponse struct{}
