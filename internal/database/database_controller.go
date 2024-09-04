package database

import (
	"go-telegram-bot/internal/database/entities"
)

type DbController struct {
	*Database
}

func NewDbController() *DbController {
	return &DbController{
		Database: NewDb(),
	}
}

func (db *DbController) AddUser(u *entities.DbTelegramUser) {
	db.Database.Db.Create(&u)
}

func (db *DbController) GetAllUsers(users *[]entities.DbTelegramUser) error {
	tx := db.Database.Db.Find(&users)
	return tx.Error
}
