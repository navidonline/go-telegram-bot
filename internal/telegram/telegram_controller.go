package telegram

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

type TelegramController struct {
	Bot *tele.Bot
}

func NewTelegramController() *TelegramController {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Panic(err)
	}

	return &TelegramController{
		Bot: b,
	}
}

func (bot *TelegramController) Start() {
	bot.Bot.Start()
}
