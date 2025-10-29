package commands

import (
	"fmt"

	"github.com/BohdanIpy/bot_256_demo/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Comander struct {
	bot      *tgbotapi.BotAPI
	service  *product.Service
	commands map[string]func(*Comander, *tgbotapi.Message)
}

func NewCommandRouter(bot_ *tgbotapi.BotAPI, service_ *product.Service) *Comander {
	return &Comander{bot: bot_, service: service_, commands: make(map[string]func(*Comander, *tgbotapi.Message))}
}

func (c *Comander) AddHandler(command string, handleFunc func(*Comander, *tgbotapi.Message)) {
	c.commands[command] = handleFunc
}

func (c *Comander) HandleUpdateMsg(update tgbotapi.Update) {
	fmt.Println(update.CallbackQuery)
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Data recv from callback query: \"%v\"", update.CallbackQuery.Data))
		c.bot.Send(msg)
		return
	}
	if !update.Message.IsCommand() {
		c.DefaultBehaviour(update.Message)
	} else {
		hFunc, found := c.commands[update.Message.Command()]
		if !found {
			c.Help(update.Message)
			return
		}
		hFunc(c, update.Message)
	}
}
