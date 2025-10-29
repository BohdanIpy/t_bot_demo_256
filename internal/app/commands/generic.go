package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) handleCommand(inputMessage *tgbotapi.Message, stringReplyFunc func() string) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, stringReplyFunc())
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
