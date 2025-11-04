package commands

import (
	"log"

	"github.com/BohdanIpy/bot_256_demo/internal/app/commands/commerce/product"
	"github.com/BohdanIpy/bot_256_demo/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
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

func (c *CommerceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "product":
		c.productCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CommerceCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "subdomain":
		c.productCommander.HandleCommand(message, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
