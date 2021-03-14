package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "appuser:password@tcp(127.0.0.1:3307)/todo_api?charset=utf8mb4&parseTime=True&loc=Local"
var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
