package commands

import (
	"github.com/BohdanIpy/bot_256_demo/internal/app/commands/commerce/product"
	"github.com/BohdanIpy/bot_256_demo/intrenal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleMessage(message *tgbotapi.Message, callbackPath path.CallbackPath)
}

type CommerceCommander struct {
	bot              *tgbotapi.BotAPI
	productCommander Commander
}

func NewCommerceCommander(bot *tgbotapi.BotAPI) *CommerceCommander {
	return &CommerceCommander{
		bot:              bot,
		productCommander: product.NewProductCommander(bot),
	}
}
