package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDatabase() (db *gorm.DB) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/ginStudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(any(err))
	}
	return db
}
