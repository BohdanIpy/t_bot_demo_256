package commands

import (
	"fmt"

	"github.com/BohdanIpy/bot_256_demo/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) List(inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, handleListCommand(c.service))

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "some data"),
		),
	)
	msg.ReplyMarkup = numericKeyboard
	if _, err := c.bot.Send(msg); err != nil {
		fmt.Println(err)
	}
	//c.handleCommand(inputMessage, func() string {
	// 	return handleListCommand(c.service)
	//})
}

func handleListCommand(productListService *product.Service) string {
	dataOutput := "List of products:"
	for _, val := range productListService.List() {
		dataOutput = dataOutput + "\n\t" + val.String()
	}

	dataOutput += "\n"
	return dataOutput
}
