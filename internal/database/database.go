package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	database *gorm.DB
}

func GetHandler() *Handler {
	dsn := "appuser:password@tcp(database:3306)/todo_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(`couldn't connect to database: ` + err.Error())
	}
	return &Handler{database: db}
}

func (h *Handler) DB() *gorm.DB {
	return h.database
}
