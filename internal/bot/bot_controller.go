package bot

import (
	"fmt"
	"go-telegram-bot/internal/database/entities"
	"go-telegram-bot/internal/entities/configs"
	"io"
	"net/http"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

type BotController struct {
	*configs.Config
}

func NewBotController(config *configs.Config) *BotController {
	return &BotController{
		Config: config,
	}
}

func (controller *BotController) Start() {
	controller.setHandlers()
	controller.Telegram.Start()
}

func (controller *BotController) setHandlers() {
	controller.handleStart()
	controller.handleGetAllUsers()
	controller.handleGetJson()
}

func (controller *BotController) handleStart() {
	controller.Telegram.Bot.Handle("/start", func(c tele.Context) error {
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

		controller.Db.AddUser(&entities.DbTelegramUser{
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
			return c.Send(controller.T("trans_received_msg", &map[string]any{"Name": getUserTitle(c.Sender()), "Id": transId}))
		}

		return c.Send(controller.getWelcomeMessage(c.Sender()))

	})
}

func (controller *BotController) getWelcomeMessage(user *tele.User) string {
	return controller.T("welcome_message", &map[string]any{"Name": getUserTitle(user)})
}

func getUserTitle(user *tele.User) string {
	return user.FirstName + " " + user.LastName
}

func (controller *BotController) handleGetAllUsers() {
	controller.Telegram.Bot.Handle("/users", func(c tele.Context) error {
		var users []entities.DbTelegramUser
		err := controller.Db.GetAllUsers(&users)
		result := ""

		if err != nil {
			return c.Send(controller.T("users_empty_msg"))
		}
		for _, user := range users {
			result += fmt.Sprintf("%v- %v	Name:%v	UserName:%v	ref:%v\n", user.ID, user.UserId, user.FirstName+" "+user.LastName, user.Username, user.RefId)
		}
		if len(result) == 0 {
			result = controller.T("users_empty_msg")
		}
		return c.Send(result)
	})
}

func (controller *BotController) handleGetJson() {
	controller.Telegram.Bot.Handle("/json", func(c tele.Context) error {
		resp, err := http.Get("https://reqres.in/api/users?page=1")
		if err != nil {
			return c.Reply("No response from request")
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return c.Reply("```json\n"+string(body)+"\n```", tele.ModeMarkdownV2)
	})
}
