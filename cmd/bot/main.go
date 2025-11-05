package main

import (
	"log"
	"os"

	routerPkg "github.com/BohdanIpy/bot_256_demo/internal/app/router"
	rp "github.com/BohdanIpy/bot_256_demo/internal/repository/commerce/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TELEGRAM_APITOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	repo := rp.NewProductRepository()
	//repo2, err := rp.NewCSVRepository("pth")
	//if err != nil {
	//	log.Fatal(err)
	//}
	routerHandler := routerPkg.NewRouter(bot, repo)
	//routerHandler2 := routerPkg.NewRouter(bot, repo2)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
