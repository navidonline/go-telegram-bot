package main

import (
	"fmt"
	"go-telegram-bot/database"
	"go-telegram-bot/lang"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Config struct {
	*database.Database
	*lang.Lang
	*tele.Bot
}

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

	cfg := Config{
		Database: database.NewDb(),
		Lang:     lang.Init(),
		Bot:      b,
	}

	handleStart(&cfg)

	handleGetAllUsers(&cfg)

	handleGetJson(&cfg)

	fmt.Println("Bot started...")
	b.Start()
}

func handleStart(cfg *Config) {
	cfg.Bot.Handle("/start", func(c tele.Context) error {
		payload := c.Message().Payload
		var refId int64
		var transId int64
		if payload != "" {
			if strings.HasPrefix(payload, "tr") {
				tr, _ := strings.CutPrefix(payload, "tr")
				transId, _ = strconv.ParseInt(tr, 0, 64)
			}
			if strings.HasPrefix(payload, "ref") {
				ref, _ := strings.CutPrefix(payload, "ref")
				refId, _ = strconv.ParseInt(ref, 0, 64)
			}
		}

		cfg.Database.Db.FirstOrCreate(&database.TelegramUser{
			UserId:       c.Sender().ID,
			RefId:        refId,
			FirstName:    c.Sender().FirstName,
			LastName:     c.Sender().LastName,
			Username:     c.Sender().Username,
			LanguageCode: c.Sender().LanguageCode,
			IsBot:        c.Sender().IsBot,
			IsPremium:    c.Sender().IsPremium,
		})

		if transId > 0 {
			return c.Send(cfg.T("trans_received_msg", &map[string]any{"Name": getUserTitle(c.Sender()), "Id": transId}))
		}

		return c.Send(getWelcomeMessage(cfg, c.Sender()))

	})
}

func getWelcomeMessage(cfg *Config, user *tele.User) string {
	return cfg.T("welcome_message", &map[string]any{"Name": getUserTitle(user)})
}

func getUserTitle(user *tele.User) string {
	return user.FirstName + " " + user.LastName
}

func handleGetAllUsers(cfg *Config) {
	cfg.Bot.Handle("/users", func(c tele.Context) error {
		var users []database.TelegramUser
		cfg.Database.Db.Find(&users)
		result := ""
		for _, user := range users {
			result += fmt.Sprintf("%v- %v	Name:%v	UserName:%v	ref:%v\n", user.ID, user.UserId, user.FirstName+" "+user.LastName, user.Username,user.RefId)
		}
		if len(result)==0{
			result=cfg.T("users_empty_msg")
		}
		return c.Send(result)
	})
}

func handleGetJson(cfg *Config) {
	cfg.Bot.Handle("/json", func(c tele.Context) error {
		resp, err := http.Get("https://reqres.in/api/users?page=1")
		if err != nil {
			return c.Reply("No response from request")
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return c.Reply("```json\n"+string(body)+"\n```", tele.ModeMarkdownV2)
	})
}
