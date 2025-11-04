package product

import (
	"encoding/json"
	"log"

	"github.com/BohdanIpy/bot_256_demo/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (p *ProductComander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products, err := p.productService.List(0, 1)
	if err != nil {
		log.Println(err)
		return
	}
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "demo",
		Subdomain:    "subdomain",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = p.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}
