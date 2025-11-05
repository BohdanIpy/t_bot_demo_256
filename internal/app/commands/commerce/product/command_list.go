package product

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/BohdanIpy/bot_256_demo/internal/app/path"
	product "github.com/BohdanIpy/bot_256_demo/internal/model/commerce"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const DEFAULT_PAGE_SIZE int64 = 4
const DEFAULT_OFFSET int64 = 0

func convertListIntoReadable(data []product.Product) string {
	var builder strings.Builder
	for _, p := range data {
		builder.WriteString(p.String())
		builder.WriteString("\n")
	}
	return builder.String()
}

func (p *ProductComander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products, err := p.productService.List(uint64(DEFAULT_OFFSET), uint64(DEFAULT_PAGE_SIZE))
	if err != nil {
		log.Println(err)
		return
	}
	outputMsgText += convertListIntoReadable(products)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: uint64(DEFAULT_OFFSET + DEFAULT_PAGE_SIZE),
		Limit:  uint64(DEFAULT_PAGE_SIZE),
	})

	callbackPath := path.CallbackPath{
		Domain:       "commerce",
		Subdomain:    "product",
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
