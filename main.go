package main

import (
	"fmt"
	"go-telegram-bot/database"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {

	db := database.NewDb()

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	handleStart(b, db)

	handleGetAllUsers(b, db)

	handleGetJson(b)

	fmt.Println("Bot started...")
	b.Start()
}

func handleStart(b *tele.Bot, db *database.Database) {
	b.Handle("/start", func(c tele.Context) error {
		db.Db.Create(&database.TelegramUser{UserId: c.Sender().ID, FirstName: c.Sender().FirstName, LastName: c.Sender().LastName, Username: c.Sender().Username})
		return c.Send(fmt.Sprint("سلام", " ", c.Sender().FirstName, " ", "خوش آمدید"))
	})
}

func handleGetAllUsers(b *tele.Bot, db *database.Database) {
	b.Handle("/users", func(c tele.Context) error {
		var users []database.TelegramUser
		db.Db.Find(&users)
		result:=""
		for _,user:=range users{
			result+=fmt.Sprintf("%v- ID:%v	UserName:%v\n",user.ID,user.UserId,user.Username)
		}
		return c.Send(result)
	})
}

func handleGetJson(b *tele.Bot) {
	b.Handle("/json", func(c tele.Context) error {
		resp, err := http.Get("https://reqres.in/api/users?page=1")
		if err != nil {
			return c.Reply("No response from request")
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return c.Reply("```json\n"+string(body)+"\n```", tele.ModeMarkdownV2)
	})
}
