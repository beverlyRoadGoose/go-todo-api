package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	database *gorm.DB
}

func GetHandler() *Handler {
	dsn := "appuser:password@tcp(127.0.0.1:3307)/todo_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Handler{database: db}
}

func (h *Handler) DB() *gorm.DB {
	return h.database
}
