package item

import (
	"github.com/google/uuid"
)

type Item struct {
	Id   uuid.UUID `json:"id" gorm:"primary_key"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
}
