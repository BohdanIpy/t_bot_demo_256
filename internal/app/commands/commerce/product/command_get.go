package product

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (p *ProductComander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := p.productService.List(0, 1)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product[0].Title,
	)

	_, err = p.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}
}
