package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct{}

var dsn = "appuser:password@tcp(127.0.0.1:3307)/todo_api?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func NewHandler() *Handler {
	return &Handler{}
}

func (dh *Handler) DB() *gorm.DB {
	return db
}