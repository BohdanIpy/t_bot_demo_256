package main

import (
	"log"
	"os"

	"github.com/BohdanIpy/bot_256_demo/internal/app/commands"
	"github.com/BohdanIpy/bot_256_demo/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

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

	comander := commands.NewCommandRouter(bot, product.NewService())
	comander.AddHandler("help", (*commands.Comander).Help)
	comander.AddHandler("list", (*commands.Comander).List)
	comander.AddHandler("help", (*commands.Comander).Help)
	comander.AddHandler("get", (*commands.Comander).GetById)

	for update := range updates {
		// log.Printf("username - %s", update.Message.Chat.UserName)
		comander.HandleUpdateMsg(update)
	}
}
