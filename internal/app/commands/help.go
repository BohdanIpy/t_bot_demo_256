package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) Help(inputMessage *tgbotapi.Message) {
	c.handleCommand(inputMessage, func() string {
		return "Allowed command:\n\thelp\n\tlist\n\tget [id]"
	})
}
