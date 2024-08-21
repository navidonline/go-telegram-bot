package main

import (
	"fmt"
	"log"
	"time"
	"os"
	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		c.Reply("salam")
		b.Send(c.Sender(), fmt.Sprint("salaaam", " ", c.Sender().FirstName))
		return nil
	})
	fmt.Println("Bot started...")
	b.Start()
}
