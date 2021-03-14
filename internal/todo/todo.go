package todo

import "github.com/google/uuid"

type Item struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
}
