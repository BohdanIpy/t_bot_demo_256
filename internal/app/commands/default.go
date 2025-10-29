package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) DefaultBehaviour(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
