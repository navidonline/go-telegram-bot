package main

import (
	"fmt"
	"go-telegram-bot/internal/bot"
	"go-telegram-bot/internal/entities/configs"
	"log"
)

func main() {

	cfg, err := configs.NewConfig()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	botController:=bot.NewBotController(cfg)
	fmt.Println("Bot started...")
	botController.Start()
}
