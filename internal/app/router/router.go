package router

import (
	"log"
	"runtime/debug"

	commerce "github.com/BohdanIpy/bot_256_demo/internal/app/commands/commerce"
	"github.com/BohdanIpy/bot_256_demo/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	bot               *tgbotapi.BotAPI
	commerceCommander Commander
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot:               bot,
		commerceCommander: commerce.NewCommerceCommander(bot),
	}
}

func (c *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "commerce":
		c.commerceCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	switch commandPath.Domain {
	case "commerce":
		c.commerceCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", commandPath.Domain)
	}
}

func (r *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	_, err := r.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
