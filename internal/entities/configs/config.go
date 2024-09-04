package configs

import (
	"go-telegram-bot/internal/database"
	"go-telegram-bot/internal/telegram"
	"go-telegram-bot/lang"
)

type Config struct {
	Db *database.DbController
	*lang.Lang
	Telegram *telegram.TelegramController
}

func NewConfig() (*Config, error) {

	return &Config{
		Db:       database.NewDbController(),
		Lang:     lang.Init(),
		Telegram: telegram.NewTelegramController(),
	}, nil
}
