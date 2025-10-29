package commands

import (
	"fmt"
	"strconv"

	"github.com/BohdanIpy/bot_256_demo/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) GetById(inputMessage *tgbotapi.Message) {
	c.handleCommand(inputMessage, func() string {
		return processGetByIdString(c.service, inputMessage.CommandArguments())
	})
}

func processGetByIdString(service *product.Service, args string) string {
	number, err := strconv.Atoi(args)
	if err != nil {
		return fmt.Sprintf("cannot convert the string in comand into an integer \"%s\"", args)
	}
	for _, val := range service.List() {
		if val.Id == number {
			return "Found the: " + val.String()
		}
	}
	return fmt.Sprintf("Not found with id: %d", number)
}
