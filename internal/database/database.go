package database

import (
	"go-telegram-bot/internal/database/entities"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func NewDb() *Database {
	return &Database{
		Db: initDb(),
	}
}

func initDb() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	db.AutoMigrate(&entities.DbTelegramUser{})

	return
}
