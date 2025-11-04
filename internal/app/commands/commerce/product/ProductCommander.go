package product

import (
	"log"

	"github.com/BohdanIpy/bot_256_demo/internal/app/path"
	"github.com/BohdanIpy/bot_256_demo/internal/service/commerce/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProductComander struct {
	bot            *tgbotapi.BotAPI
	productService product.ProductService
}

func NewProductCommander(bot *tgbotapi.BotAPI) *ProductComander {
	panic("TODO")
	/*
		service := product.NewDummyProductService()
		return &ProductComander{
			bot:            bot,
			productService: service,
		}
	*/
}

func (p *ProductComander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		p.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (p *ProductComander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		p.Help(msg)
	case "list":
		p.List(msg)
	case "get":
		p.Get(msg)
	default:
		p.Default(msg)
	}
}
