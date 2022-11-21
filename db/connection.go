package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GetConnection() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:8889)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		log.Println("Database Connected")
		return db
	}
}
