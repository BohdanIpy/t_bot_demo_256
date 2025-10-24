package main

import (
	"log"
	"os"

	"github.com/BohdanIpy/bot_256_demo/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Allowed command:\n\tsayhi\n\tstatus")
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func defaultBehaviour(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func handleCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, stringReplyFunc func() string) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, stringReplyFunc())
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func handleListCommand(productListService *product.Service) string {
	dataOutput := "List of products:"
	for _, val := range productListService.List() {
		dataOutput = dataOutput + "\n\t" + val.String()
	}
	dataOutput += "\n"
	return dataOutput
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		log.Printf("username - %s", update.Message.Chat.UserName)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				handleCommand(bot, update.Message, func() string {
					return handleListCommand(productService)
				})
			default:
				helpCommand(bot, update.Message)
				return
			}
		} else {
			defaultBehaviour(bot, update.Message)
		}
	}
}
