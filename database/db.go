package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TelegramUser struct {
	gorm.Model
	UserId       int64  `json:"id" gorm:"primaryKey"`
	RefId        int64  `json:"ref_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsBot        bool   `json:"is_bot"`
	IsPremium    bool   `json:"is_premium"`
}

type Database struct {
	Db *gorm.DB
}

func NewDb() *Database {
	return &Database{
		Db: initDb(),
	}
}

func initDb() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TelegramUser{})

	return
}
