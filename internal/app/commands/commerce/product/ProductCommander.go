package product

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type ProductCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

func NewProductCommander(bot *tgbotapi.BotAPI, service *service.ProductService) *ProductCommander {
	return &ProductComandImpl{}
}

type ProductComandImpl struct {
}
