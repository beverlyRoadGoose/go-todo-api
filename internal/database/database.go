package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo-api/pkg/config"
)

type Handler struct {
	database *gorm.DB
}

func GetHandler() *Handler {
	dbConf := config.Conf.Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/todo_api?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(`couldn't connect to database: ` + err.Error())
	}
	return &Handler{database: db}
}

func (h *Handler) DB() *gorm.DB {
	return h.database
}
