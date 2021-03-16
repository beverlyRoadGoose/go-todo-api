package item

import (
	"github.com/google/uuid"
	"todo-api/internal/database"
)

type Item struct {
	Id   uuid.UUID `json:"id" gorm:"primary_key"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
}

func init() {
	dh := database.GetHandler()
	dh.DB().AutoMigrate(&Item{})
}
